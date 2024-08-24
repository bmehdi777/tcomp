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


func (t *Tmux) NewSession(name string) error {
	return t.execCmd("new-session", "-ds", name)
}

func (t *Tmux) NewWindow() error {
	return t.execCmd("new-window")
}

func (t *Tmux) NewWindowWithCommand(programName string) error {
	return t.execCmd("new-window", programName)
}

func (t *Tmux) SplitHorizontal() error {
	return t.execCmd("split-window", "-h")
}

func (t *Tmux) SplitHorizontalWithCommand(programName string) error {
	return t.execCmd("split-window", "-h", programName)
}

func (t *Tmux) SplitVertical() error {
	return t.execCmd("split-window", "-v")
}

func (t *Tmux) SplitVerticalWithCommand(programName string) error {
	return t.execCmd("split-window", "-v", programName)
}
