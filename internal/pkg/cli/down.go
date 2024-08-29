package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/bmehdi777/tcomp/internal/pkg/tmux"
	"github.com/bmehdi777/tcomp/internal/pkg/workspace"
	"github.com/spf13/cobra"
)

func newCmdDown() *cobra.Command {
	downCmd := cobra.Command{
		Use:   "down [WORKSPACE]",
		Short: "Stop your tmux environment",
		Run:   handlerDown,
		Args:  cobra.RangeArgs(0, 1),
	}

	downCmd.PersistentFlags().StringP("file", "f", "", "Path to the workspace's file")

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
		pathFile, err = workspace.GetFileWorkspacePath(args[0], config)
		checkError(err)
	}

	ws, err := workspace.ReadWorkspace(pathFile)
	checkError(err)

	err = ws.StopTmuxEnv(config)
	checkError(err)

	fmt.Printf("Session %v is down.", ws.Session)
}
