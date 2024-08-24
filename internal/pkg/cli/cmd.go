package cli

import (
	"fmt"
	"os"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
	"github.com/spf13/cobra"
)

func newCmdRoot() *cobra.Command {
	rootCmd := cobra.Command {
		Use: "tcompose",
		Short: "Compose your tmux environment",
	}

	_, err := tmux.NewConfig()
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	repo, err := tmux.ReadRepository("./examples/basic.yml")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	fmt.Println("repo", repo)

	return &rootCmd
}

func Execute() error {
	return newCmdRoot().Execute()
}
