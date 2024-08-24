package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newCmdRoot() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "tcompose",
		Short: "Compose your tmux environment",
	}

	rootCmd.AddCommand(newCmdUp())

	return &rootCmd
}

func Execute() error {
	return newCmdRoot().Execute()
}

func checkGenericError(err error) {
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
}
