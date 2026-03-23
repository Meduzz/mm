package cmd

import (
	"os/exec"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/mm/pkg/config"
	"github.com/spf13/cobra"
)

func serve() *model.Command {
	return commando.CommandBuilder("serve", func(cb builder.CommandBuilder) {
		cb.Description("Starts llama.cpp server.")
		cb.Handler(serveHandler)
		cb.Flag(flags.IntFlag("port", 8080, "The port to run the server on."))
	})
}

func serveHandler(cmd *cobra.Command, args []string) error {
	port, _ := cmd.Flags().GetInt("port")

	cfg, err := config.LoadConfig()

	if err != nil {
		return err
	}

	cmdPath := "llama-server"
	cmdArgs := cfg.Attributes

	runtimeFile := &config.Runtime{
		Port: port,
	}

	err = config.StoreRuntime(runtimeFile)

	if err != nil {
		return err
	}

	serverCmd := exec.Command(cmdPath, cmdArgs...)

	if err := serverCmd.Start(); err != nil {
		return err
	}

	runtimeFile.PID = serverCmd.Process.Pid

	err = config.StoreRuntime(runtimeFile)
	if err != nil {
		serverCmd.Process.Kill()
		return err
	}

	return nil
}
