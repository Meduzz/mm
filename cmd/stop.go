package cmd

import (
	"os"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/mm/pkg/config"
	"github.com/spf13/cobra"
)

func stop() *model.Command {
	return commando.CommandBuilder("stop", func(cb builder.CommandBuilder) {
		cb.Description("Stop the llama.cpp server.")
		cb.Handler(stopHandler)
	})
}

func stopHandler(cmd *cobra.Command, args []string) error {
	runtime, err := config.LoadRuntime()

	if err != nil {
		return err
	}

	process, err := os.FindProcess(runtime.PID)

	if err != nil {
		return err
	}

	err = process.Kill()

	if err != nil {
		return err
	}

	err = config.RemoveRuntime()

	if err != nil {
		return err
	}

	return nil
}
