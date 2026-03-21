package cmd

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func ls() *model.Command {
	return commando.CommandBuilder("ls", func(cb builder.CommandBuilder) {
		cb.Description("List models available in llama.cpp.")
		cb.Handler(lsHandler)
	})
}

func lsHandler(cmd *cobra.Command, args []string) error {
	return nil
}
