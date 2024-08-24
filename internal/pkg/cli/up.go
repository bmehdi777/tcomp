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

	fmt.Println("repo", repo)

	tmux := tmux.Tmux{Config: config}
	err = tmux.NewWindow("TEssssssssssssst 2").Execute("ls", "zsh")
	checkGenericError(err)

	err = tmux.NewSplitPaneVertical().Execute("echo hello", "zsh")
	checkGenericError(err)

	err = tmux.NewSplitPaneHorizontal().Execute()
	checkGenericError(err)
}
