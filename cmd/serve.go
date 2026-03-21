package cmd

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/registry"
	"github.com/spf13/cobra"
)

func init() {
	serve := commando.Command("serve", serveHandler)
	registry.RegisterCommand(serve)
}

func serveHandler(cmd *cobra.Command, args []string) error {
	// TODO implement
	return nil
}
