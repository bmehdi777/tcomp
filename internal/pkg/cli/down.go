package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/repository"
	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
	"github.com/spf13/cobra"
)

func newCmdDown() *cobra.Command {
	downCmd := cobra.Command{
		Use:   "down [REPOSITORY]",
		Short: "Stop your tmux environment",
		Run:   handlerDown,
		Args:  cobra.RangeArgs(0, 1),
	}

	downCmd.PersistentFlags().StringP("file", "f", "", "Path to the repository's file")

	return &downCmd
}

func handlerDown(cmd *cobra.Command, args []string) {
	config, err := tmux.NewConfig()
	checkError(err)

	var pathFile string
	if len(args) <= 0 {
		pathFile, err = cmd.PersistentFlags().GetString("file")
		checkError(err)
		if pathFile == "" {
			cmd.Usage()
			os.Exit(1)
		}

		if _, err := os.Stat(pathFile); errors.Is(err, os.ErrNotExist) {
			fmt.Println("Filepath is not valid.")
			os.Exit(1)
		}
	} else {
		pathFile, err = repository.GetFileRepoPath(args[0], config)
		checkError(err)
	}

	repo, err := repository.ReadRepository(pathFile)
	checkError(err)

	err = repo.StopTmuxEnv(config)
	checkError(err)

	fmt.Printf("Session %v is down.", repo.Session)
}
