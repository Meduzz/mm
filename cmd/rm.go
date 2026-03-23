package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/mm/pkg/config"
	"github.com/Meduzz/mm/pkg/openai"
	"github.com/spf13/cobra"
)

func rm() *model.Command {
	return commando.CommandBuilder("rm", func(cb builder.CommandBuilder) {
		cb.Description("Remove the specified model from llama directory.")
		cb.Handler(rmHandler)
		cb.Flag(flags.StringFlag("model", "", "The name of the model to remove."))
	})
}

func deleteModel(match string) error {
	if err := os.Remove(match); err != nil {
		return err
	}

	return nil
}

func rmHandler(cmd *cobra.Command, args []string) error {
	modelName, err := cmd.Flags().GetString("model")

	if err != nil {
		return err
	}

	if modelName == "" {
		return fmt.Errorf("model name is required")
	}

	models, err := openai.LoadModels()

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Is the server runnig?")
		}
		return err
	}

	targetModel := slice.Head(slice.Filter(models, func(m *openai.Model) bool {
		return m.Name == modelName
	}))

	if targetModel == nil {
		fmt.Printf("Model %s does not exist.\n", modelName)
		return nil
	}

	if targetModel.Path != "" {
		dir := filepath.Dir(targetModel.Path)
		matches, err := filepath.Glob(filepath.Join(dir, modelName+"*"))
		if err != nil {
			return err
		}

		for _, match := range matches {
			if err := deleteModel(match); err != nil {
				return err
			}
		}
	} else {
		cfg, err := config.LoadConfig()

		if err != nil {
			return err
		}

		// TODO not all models will contain a slash
		parts := strings.Split(modelName, "/")

		matches, err := filepath.Glob(filepath.Join(cfg.Models, "*"+parts[1]+"*.*"))

		if err != nil {
			return err
		}

		for _, match := range matches {
			if err := deleteModel(match); err != nil {
				return err
			}
		}
		fmt.Printf("%s removed successfully.\n", modelName)
	}

	return nil
}
