package cli

import (
	"fmt"
	"strings"

	"github.com/bmehdi777/tcomp/internal/pkg/tmux"
	"github.com/bmehdi777/tcomp/internal/pkg/workspace"
	"github.com/spf13/cobra"
)

func newCmdList() *cobra.Command {
	listCmd := cobra.Command{
		Use:     "ls",
		Short:   "List every workspace file you have stored",
		Run:     handlerList,
		Aliases: []string{"list", "see"},
	}

	return &listCmd
}

func handlerList(cmd *cobra.Command, args []string) {
	config, err := tmux.NewConfig()
	checkError(err)

	listName, err := workspace.ListWorkspace(config)
	if len(listName) > 0 {
		fmt.Printf("List of workspace files :\n%v\n", strings.Join(listName, "\n"))
	} else {
		fmt.Printf("No workspace.\nSee `tcomp new <WORKSPACE>` to create a new one.\n")
	}
}
