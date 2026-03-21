package cmd

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func rm() *model.Command {
	return commando.CommandBuilder("rm", func(cb builder.CommandBuilder) {
		cb.Description("Remove the specified model from llama directory.")
		cb.Handler(rmHandler)
		cb.Flag(flags.StringFlag("model", "", "The name of the model to remove."))
	})
}

func rmHandler(cmd *cobra.Command, args []string) error {
	// TODO load model flag.
	// TODO make sure it's set.
	return nil
}
