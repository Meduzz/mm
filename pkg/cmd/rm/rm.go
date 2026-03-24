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

type Rm struct{}

func New() *Rm {
	return &Rm{}
}

func (r *Rm) DeleteModel(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func (r *Rm) getGlobPattern(name string) string {
	if strings.Contains(name, "/") {
		parts := strings.Split(name, "/")
		name = parts[1]
	}

	return "*" + name + "*.*"
}

func (r *Rm) RemoveModel(modelName string) error {
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
	filePattern := r.getGlobPattern(modelName)

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
		if err := r.DeleteModel(match); err != nil {
			return err
		}
	}

	fmt.Printf("%s removed successfully\n", modelName)
	return nil
}
