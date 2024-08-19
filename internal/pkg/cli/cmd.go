package cli

import "github.com/spf13/cobra"

func newCmdRoot() *cobra.Command {
	rootCmd := cobra.Command {
		Use: "tcompose",
		Short: "Compose your tmux environment",
	}

	return &rootCmd
}

func Execute() error {
	return newCmdRoot().Execute()
}
