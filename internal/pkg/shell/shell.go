package shell

import (
	_ "embed"
	"fmt"
)

//go:embed scripts/completion.zsh
var zshScript string

func ShowZshScript() {
	fmt.Printf(zshScript)
}

