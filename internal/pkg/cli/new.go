package cli

import (
	"fmt"

	"github.com/bmehdi777/tcomp/internal/pkg/tmux"
	"github.com/bmehdi777/tcomp/internal/pkg/workspace"
	"github.com/spf13/cobra"
)

func newCmdNew() *cobra.Command {
	newCmd := cobra.Command{
		Use:     "new <NAME>",
		Short:   "Create a new workspace file",
		Run:     handlerNew,
		Args:    cobra.RangeArgs(1, 1),
		Aliases: []string{"create"},
	}

	return &newCmd
}

func handlerNew(cmd *cobra.Command, args []string) {
	config, err := tmux.NewConfig()
	checkError(err)

	filePath, err := workspace.CreateNewWorkspaceFile(args[0], config)
	checkError(err)

	workspace.OpenWorkspaceFileWithEditor(filePath, config)
	fmt.Printf("Successfully created %v workspace !\n", args[0])
}
