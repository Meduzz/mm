package cmd

import (
	"fmt"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/mm/pkg/cmd/rm"
	"github.com/spf13/cobra"
)

func removeCommand() *model.Command {
	return commando.CommandBuilder("rm", func(cb builder.CommandBuilder) {
		cb.Description("Remove the specified model from llama directory.")
		cb.Handler(rmHandler)
		cb.Flag(flags.StringFlag("model", "", "The name of the model to remove."))
	})
}

func rmHandler(cmd *cobra.Command, args []string) error {
	modelName, err := cmd.Flags().GetString("model")

	if err != nil {
		return err
	}

	if modelName == "" {
		return fmt.Errorf("model name is required")
	}

	return rm.RemoveModel(modelName)
}
