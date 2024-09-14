package cli

import (
	"fmt"

	_ "embed"

	"github.com/bmehdi777/tcomp/internal/pkg/shell"
	"github.com/spf13/cobra"
)

func newCmdInit() *cobra.Command {
	initCmd := cobra.Command{
		Use:   "init <SHELL>",
		Short: "Print tcomp's shell init script",
		Run:   handlerInit,
		Args:  cobra.RangeArgs(1, 1),
	}

	return &initCmd
}

func handlerInit(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "zsh":
		shell.ShowZshScript()
	case "fish":
	case "bash":
	default:
		fmt.Printf("%v's shell is not recognized or supported. Current shell supported :\n- zsh\n", args[0])
	}
}
