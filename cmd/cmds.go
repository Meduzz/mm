package cmd

import "github.com/Meduzz/commando/registry"

func init() {
	registry.RegisterCommand(serve())
	registry.RegisterCommand(ls())
	registry.RegisterCommand(rm())
	registry.RegisterCommand(unload())
}
