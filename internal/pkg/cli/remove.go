package cli

import (
	"github.com/bmehdi777/tcomp/internal/pkg/tmux"
	"github.com/bmehdi777/tcomp/internal/pkg/workspace"
	"github.com/spf13/cobra"
)

func newCmdRemove() *cobra.Command {
	removeCmd := cobra.Command{
		Use:     "remove <NAME>",
		Short:   "Remove a workspace file",
		Run:     handlerRm,
		Args:    cobra.MinimumNArgs(1),
		Aliases: []string{"rm", "delete", "del"},
	}

	return &removeCmd
}

func handlerRm(cmd *cobra.Command, args []string) {
	config, err := tmux.NewConfig()
	checkError(err)

	for _, file := range args {
		err = workspace.RemoveWorkspaceFile(file, config)
		checkError(err)
	}
}
