package tmux

import (
	"fmt"
	"os/exec"
	"slices"
)

type Tmux struct {
	Conf *Config
}

func (t *Tmux) execCmd(param ...string) error {
	fullParam := slices.Insert(param, 0, "-S")
	fullParam = slices.Insert(fullParam, 1, t.Conf.TmuxSocketPath)
	fmt.Println("Full param", fullParam)

	cmd := exec.Command("tmux", fullParam...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (t *Tmux) NewSession(name string) {
	t.execCmd("new-session", "-ds", name)
}
