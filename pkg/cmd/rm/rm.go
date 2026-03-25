package rm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/mm/pkg/config"
	"github.com/Meduzz/mm/pkg/openai"
)

func deleteModel(path string) error {
	err := os.Remove(path)

	if err != nil {
		return err
	}

	return nil
}

func getGlobPattern(name string) string {
	if strings.Contains(name, "/") {
		parts := strings.Split(name, "/")
		name = parts[1]
	}

	return "*" + name + "*.*"
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

	// simplify filename remove /
	filePattern := getGlobPattern(modelName)

	// start with dir for model from server (can be empty)
	dir := targetModel.Path

	if dir == "" {
		// instead load it from config
		cfg, err := config.LoadConfig()

		if err != nil {
			return err
		}

		dir = cfg.Models
	}

	// find files matching filePattern
	matches, err := filepath.Glob(filepath.Join(dir, filePattern))

	if err != nil {
		return err
	}

	// remove matches
	for _, match := range matches {
		err := deleteModel(match)

		if err != nil {
			return err
		}
	}

	fmt.Printf("%s removed successfully\n", modelName)
	return nil
}
