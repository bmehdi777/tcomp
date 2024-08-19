package cli

import "github.com/spf13/cobra"

func NewCmdRoot() *cobra.Command {
	rootCmd := cobra.Command {
		Use: "tmuxcompose",
		Short: "Compose your tmux environment",
	}

	return &rootCmd
}
