package cmd

import (
	"fmt"
	"os"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/mm/pkg/openai"
	"github.com/spf13/cobra"
)

func ls() *model.Command {
	return commando.CommandBuilder("ls", func(cb builder.CommandBuilder) {
		cb.Description("List models available in llama.cpp.")
		cb.Handler(lsHandler)
	})
}

func lsHandler(cmd *cobra.Command, args []string) error {
	models, err := openai.LoadModels()

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Is the server runnig?")
		}

		return err
	}

	fmt.Println("Models:")
	slice.ForEach(models, func(m *openai.Model) {
		fmt.Printf("- %s (loaded:%v)\n", m.Name, m.Loaded)
	})

	return nil
}
