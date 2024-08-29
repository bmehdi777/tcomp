package cli

import (
	"fmt"

	"github.com/bmehdi777/tcomp/internal/pkg/version"
	"github.com/spf13/cobra"
)

func newCmdVersion() *cobra.Command {
	versionCmd := cobra.Command{
		Use: "version",
		Run: handlerVersion,
	}

	return &versionCmd
}

func handlerVersion(cmd *cobra.Command, args []string) {
	versionInfo := version.Get()
	fmt.Println("Version : ", versionInfo.Version)
}
