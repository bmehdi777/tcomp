package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newCmdRoot() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "tcomp",
		Short: "Compose your tmux environment",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(newCmdUp())
	rootCmd.AddCommand(newCmdDown())
	rootCmd.AddCommand(newCmdList())
	rootCmd.AddCommand(newCmdNew())
	rootCmd.AddCommand(newCmdVersion())
	rootCmd.AddCommand(newCmdInit())
	rootCmd.AddCommand(newCmdRemove())

	return &rootCmd
}

func Execute() error {
	return newCmdRoot().Execute()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
}
