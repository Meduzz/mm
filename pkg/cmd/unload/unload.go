package unloadcmd

import (
	"fmt"
	"os"

	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/mm/pkg/openai"
)

func UnloadModel(name string) error {
	models, err := openai.LoadModels()

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Is the server runnig?")
		}

		return err
	}

	if name != "" {
		models = slice.Filter(models, func(m *openai.Model) bool {
			return m.Name == name
		})
	}

	for _, m := range models {
		if m.Loaded {
			err := openai.UnloadModel(m.Name)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
