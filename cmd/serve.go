package cmd

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func serve() *model.Command {
	return commando.CommandBuilder("serve", func(cb builder.CommandBuilder) {
		cb.Description("Starts llama.cpp server.")
		cb.Handler(serveHandler)
	})
}

func serveHandler(cmd *cobra.Command, args []string) error {
	// TODO create json-file with port that server was started on, also add pid to the file.
	return nil
}
