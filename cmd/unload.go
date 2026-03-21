package cmd

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func unload() *model.Command {
	return commando.CommandBuilder("unload", func(cb builder.CommandBuilder) {
		cb.Description("Unload the specified model or all models that are loaded.")
		cb.Handler(unloadHandler)
		cb.Flag(flags.StringFlag("model", "", "The name of the model to be unloaded."))
	})
}

func unloadHandler(cmd *cobra.Command, args []string) error {
	// TODO read model flag.
	return nil
}
