package cli

import (
	"fmt"
	"os"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
	"github.com/spf13/cobra"
)

func newCmdRoot() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "tcompose",
		Short: "Compose your tmux environment",
	}

	config, err := tmux.NewConfig()
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	repo, err := tmux.ReadRepository("./examples/basic.yml")
	checkGenericError(err)

	fmt.Println("repo", repo)

	tmux := tmux.Tmux{Config: config}
	err = tmux.NewWindow("TEssssssssssssst").SetCWD("/").Execute()
	checkGenericError(err)


	return &rootCmd
}

func Execute() error {
	return newCmdRoot().Execute()
}

func checkGenericError(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
}
