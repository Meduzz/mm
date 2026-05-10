package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/mm/pkg/config"
	"github.com/Meduzz/mm/pkg/openai"
	"github.com/spf13/cobra"
)

func serve() *model.Command {
	return commando.CommandBuilder("serve", func(cb builder.CommandBuilder) {
		cb.Description("Starts llama.cpp server.")
		cb.Handler(serveHandler)
		cb.Flag(flags.IntFlag("port", 8080, "The port to run the server on."))
		cb.Flag(flags.IntFlag("context", 32, "Start server with this max context length in GB"))
	})
}

func serveHandler(cmd *cobra.Command, args []string) error {
	port, _ := cmd.Flags().GetInt("port")
	context, _ := cmd.Flags().GetInt("context")

	cfg, err := config.LoadConfig(port, context*1024)

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

	fmt.Println(serverCmd.String())

	err = serverCmd.Start()

	if err != nil {
		return err
	}

	time.Sleep(250 * time.Millisecond)

	err = openai.Health()

	if err != nil {
		// the server most likely did not start
		ps, err := serverCmd.Process.Wait()

		if err != nil {
			return err
		}

		return fmt.Errorf("%s exited with %d", serverCmd.String(), ps.ExitCode())
	}

	runtimeFile.PID = serverCmd.Process.Pid

	err = config.StoreRuntime(runtimeFile)

	if err != nil {
		serverCmd.Process.Kill()
		return err
	}

	fmt.Printf("llama-server started on port: %d.\n", port)

	return nil
}
