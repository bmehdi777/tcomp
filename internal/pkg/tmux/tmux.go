package tmux

import (
	"os/exec"
	"strings"
)

type Tmux struct {
	Config *Config
	Envs   *map[string]string
}

type TmuxCommand struct {
	conf      *Config
	command   string
	params    []string
	globalEnv *map[string]string
	localEnv  *map[string]string
}

func (t *Tmux) NewSession(name string) *TmuxCommand {
	cmd := TmuxCommand{
		conf:      t.Config,
		command:   "new-session",
		params:    []string{"-ds", name},
		globalEnv: t.Envs,
	}

	return &cmd
}

func (t *Tmux) NewWindow(sessionName string, name string) *TmuxCommand {
	cmd := TmuxCommand{
		conf:      t.Config,
		command:   "new-window",
		params:    []string{"-n", name, "-t", sessionName},
		globalEnv: t.Envs,
	}

	return &cmd
}

func (t *Tmux) RenameWindow(sessionName string, currentName string, newName string) *TmuxCommand {
	cmd := TmuxCommand{
		conf:      t.Config,
		command:   "rename-window",
		params:    []string{"-t", sessionName + ":" + currentName, newName},
		globalEnv: t.Envs,
	}

	return &cmd
}

func (t *Tmux) NewSplitPaneHorizontal(sessionName string, windowName string) *TmuxCommand {
	cmd := TmuxCommand{
		conf:      t.Config,
		command:   "split-window",
		params:    []string{"-h", "-t", sessionName},
		globalEnv: t.Envs,
	}

	return &cmd
}

func (t *Tmux) NewSplitPaneVertical(sessionName string, windowName string) *TmuxCommand {
	cmd := TmuxCommand{
		conf:      t.Config,
		command:   "split-window",
		params:    []string{"-v", "-t", sessionName + ":" + windowName},
		globalEnv: t.Envs,
	}

	return &cmd
}

func (tc *TmuxCommand) SetCWD(path string) *TmuxCommand {
	tc.params = append(tc.params, "-c", path)

	return tc
}

func (tc *TmuxCommand) SetEnv(envs *map[string]string) *TmuxCommand {
	tc.localEnv = envs

	return tc
}

func (tc *TmuxCommand) Execute(programs ...string) error {
	fullParam := append([]string{"-S", tc.conf.TmuxSocketPath, tc.command}, tc.params...)

	for _, program := range programs {
		program = strings.ReplaceAll(program, "\"", "\\\"")
	}

	fullProgram := ""
	if len(programs) > 0 {
		fullProgram = strings.Join(programs, "; ")
		fullParam = append(fullParam, fullProgram)
	}

	var envs []string
	if tc.localEnv != nil && tc.globalEnv != nil {
		for key, value := range *tc.globalEnv {
			envs = append(envs, key+"="+value)
		}
		for key, value := range *tc.localEnv {
			envs = append(envs, key+"="+value)
		}
	}

	cmd := exec.Command("tmux", fullParam...)
	if len(envs) > 0 {
		cmd.Env = envs
	}
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
