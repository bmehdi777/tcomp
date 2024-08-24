package cli

import (
	"fmt"
	"strings"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/repository"
	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
	"github.com/spf13/cobra"
)

func newCmdList() *cobra.Command {
	listCmd := cobra.Command{
		Use:     "ls",
		Short:   "List every repository file you have stored",
		Run:     handlerList,
		Aliases: []string{"list", "see"},
	}

	return &listCmd
}

func handlerList(cmd *cobra.Command, args []string) {
	config, err := tmux.NewConfig()
	checkError(err)

	listName, err := repository.ListRepository(config)
	fmt.Printf("List of repository files :\n%v\n", strings.Join(listName, "\n"))
}
