package cli

import (
	"fmt"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
	"github.com/spf13/cobra"
)

func newCmdUp() *cobra.Command {
	upCmd := cobra.Command{
		Use:   "up",
		Short: "Start your tmux environment",
		Run:   handlerUp,
	}

	return &upCmd
}

func handlerUp(cmd *cobra.Command, args []string) {
	config, err := tmux.NewConfig()
	checkGenericError(err)

	repo, err := tmux.ReadRepository("./examples/basic.yml")
	checkGenericError(err)

	err = repo.ParseToTmux(config)
	checkGenericError(err)

	fmt.Println("Session is up !")
}
