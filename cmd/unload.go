package cmd

import (
	"fmt"
	"os"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/mm/pkg/openai"
	"github.com/spf13/cobra"
)

func unload() *model.Command {
	return commando.CommandBuilder("unload", func(cb builder.CommandBuilder) {
		cb.Description("Unload the specified model from llama.cpp.")
		cb.Handler(unloadHandler)
		cb.Flag(flags.StringFlag("model", "", "The name of the model to be unloaded."))
	})
}

func unloadHandler(cmd *cobra.Command, args []string) error {
	modelName, _ := cmd.Flags().GetString("model")

	models, err := openai.LoadModels()

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Is the server runnig?")
		}

		return err
	}

	if modelName == "" {
		for _, m := range models {
			if m.Loaded {
				err := openai.UnloadModel(m.Name)

				if err != nil {
					return err
				}
			}
		}
	} else {
		err := openai.UnloadModel(modelName)

		if err != nil {
			return err
		}
	}

	return nil
}
