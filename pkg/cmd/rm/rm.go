package rm

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/mm/pkg/config"
	"github.com/Meduzz/mm/pkg/openai"
)

func deleteModel(path string) error {
	return os.RemoveAll(path)
}

func getGlobPattern(name string) string {
	if strings.Contains(name, "/") {
		parts := strings.Split(name, "/")
		name = parts[len(parts)-1]
	}

	return "*" + name + "*"
}

func RemoveModel(modelName string) error {
	// list models
	models, err := openai.LoadModels()

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Is the server running?")
		}
		return err
	}

	// find our model by name
	targetModel := slice.Head(slice.Filter(models, func(m *openai.Model) bool {
		return m.Name == modelName
	}))

	if targetModel == nil {
		return fmt.Errorf("model %s does not exist", modelName)
	}

	// start with dir for model from server (can be empty)
	dir := targetModel.Path

	if dir == "" {
		// instead load it from config
		cfg, err := config.LoadConfig(0, 0)

		if err != nil {
			return err
		}

		dir = cfg.Models
	}

	// build the base name
	filePattern := fmt.Sprintf("models/%s", modelName)
	// replace slash by --
	filePattern = strings.ReplaceAll(filePattern, "/", "--")
	// remove tag-name
	split := strings.Split(filePattern, ":")
	// add dir.
	filePattern = path.Join(dir, split[0])

	if err != nil {
		return err
	}

	// remove stuff.
	err = deleteModel(filePattern)

	if err != nil {
		return err
	}

	fmt.Printf("%s removed successfully\n", modelName)
	return nil
}
