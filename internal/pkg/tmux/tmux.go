package tmux

import (
	"os/exec"
	"strings"
)

type Tmux struct {
	Config *Config
}

type TmuxCommand struct {
	conf    *Config
	command string
	params  []string
}

func (t *Tmux) NewSession(name string) *TmuxCommand {
	cmd := TmuxCommand{
		conf:    t.Config,
		command: "new-session",
		params:  []string{"-ds", name},
	}

	return &cmd
}

func (t *Tmux) NewWindow(name string) *TmuxCommand {
	name = strings.ReplaceAll(name, "'", "\\'")
	name = "'" + name + "'"

	cmd := TmuxCommand{
		conf:    t.Config,
		command: "new-window",
		params:  []string{"-n", name},
	}

	return &cmd
}

func (t *Tmux) NewSplitPaneHorizontal() *TmuxCommand {
	cmd := TmuxCommand{
		conf:    t.Config,
		command: "split-window",
		params:  []string{"-h"},
	}

	return &cmd
}

func (t *Tmux) NewSplitPaneVertical() *TmuxCommand {
	cmd := TmuxCommand{
		conf:    t.Config,
		command: "split-window",
		params:  []string{"-v"},
	}

	return &cmd
}

func (tc *TmuxCommand) SetCWD(path string) *TmuxCommand {
	tc.params = append(tc.params, "-c", path)

	return tc
}

func (tc *TmuxCommand) SetEnv(envs map[string]string) *TmuxCommand {
	return tc
}

func (tc *TmuxCommand) Execute(programs ...string) error {
	for _, program := range programs {
		program = strings.ReplaceAll(program, "'", "\\'")
	}

	fullParam := append([]string{"-S", tc.conf.TmuxSocketPath, tc.command}, tc.params...)

	fullProgram := ""
	if len(programs) > 0 {
		fullProgram = "'" + strings.Join(programs, ";") + "'"
		fullParam = append(fullParam, fullProgram)
	}

	cmd := exec.Command("tmux", fullParam...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
