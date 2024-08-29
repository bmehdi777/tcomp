package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/bmehdi777/tcomp/internal/pkg/tmux"
	"github.com/bmehdi777/tcomp/internal/pkg/workspace"
	"github.com/spf13/cobra"
)

func newCmdUp() *cobra.Command {
	upCmd := cobra.Command{
		Use:   "up [REPOSITORY]",
		Short: "Start your tmux environment",
		Run:   handlerUp,
	}

	upCmd.PersistentFlags().StringP("file", "f", "", "Path to the repository's file")

	return &upCmd
}

func handlerUp(cmd *cobra.Command, args []string) {
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
		pathFile, err = workspace.GetFileWorkspacePath(args[0], config)
		checkError(err)
	}

	ws, err := workspace.ReadWorkspace(pathFile)

	err = ws.StartTmuxEnv(config)
	checkError(err)

	fmt.Printf("Session %v is up !", ws.Session)
}
