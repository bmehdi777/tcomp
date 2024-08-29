package cli

import (
	"fmt"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/repository"
	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
	"github.com/spf13/cobra"
)

func newCmdNew() *cobra.Command {
	newCmd := cobra.Command{
		Use:     "new <NAME>",
		Short:   "Create a new repository file",
		Run:     handlerNew,
		Args:    cobra.RangeArgs(1, 1),
		Aliases: []string{"create"},
	}

	return &newCmd
}

func handlerNew(cmd *cobra.Command, args []string) {
	config, err := tmux.NewConfig()
	checkError(err)

	filePath, err := repository.CreateNewRepoFile(args[0], config)
	checkError(err)

	repository.OpenRepoFileWithEditor(filePath, config)
	fmt.Printf("Successfully created %v repository !", args[0])
}
