package workspace

import (
	"errors"
	"fmt"
	"os"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
	"gopkg.in/yaml.v3"
)

type Workspace struct {
	Session    string            `yaml:"session"`          // mandatory
	Before     []string          `yaml:"before,omitempty"` // optional
	Stop       []string          `yaml:"stop,omitempty"`   // optional
	Env        map[string]string `yaml:"env,omitempty"`    // optional
	Follow     bool              `yaml:"follow"`           // optional
	CwdSession string            `yaml:"cwd"`              // optional
	Windows    []WorkspaceWindow      `yaml:"windows"`          // mandatory
}

type WorkspaceWindow struct {
	Name      string     `yaml:"name"`               // mandatory
	Cwd       string     `yaml:"cwd"`                // optional - get current working dir
	Commands  []string   `yaml:"commands,omitempty"` // optional - do nothing
	Panes     []WorkspacePane `yaml:"panes,omitempty"`    // optional
	KeepAlive bool       `yaml:"keep_alive"`         //optional
}

type WorkspacePane struct {
	Type      WorkspacePaneType `yaml:"type,omitempty"`     // mandatory
	Cwd       string       `yaml:"cwd"`                // optional
	Commands  []string     `yaml:"commands,omitempty"` // optional
	KeepAlive bool         `yaml:"keep_alive"`         //optional
}

type WorkspacePaneType string

const (
	Horizontal WorkspacePaneType = "horizontal"
	Vertical   WorkspacePaneType = "vertical"
)

func ReadWorkspace(pathfile string) (Workspace, error) {
	ws := Workspace{}
	data, err := os.ReadFile(pathfile)
	if err != nil {
		return ws, err
	}

	err = yaml.Unmarshal(data, &ws)
	if err != nil {
		fmt.Println("Err :%w")
		return ws, err
	}

	err = ws.verifyWorkspace()
	if err != nil {
		return Workspace{}, err
	}

	return ws, nil
}

func (ws *Workspace) verifyWorkspace() error {
	if ws.Session == "" {
		return errors.New(fmt.Sprintf("`session` is missing."))
	}
	if ws.Windows == nil {
		return errors.New(fmt.Sprintf("`windows` is missing."))
	}

	for indexWindow, window := range ws.Windows {
		if window.Name == "" {
			return errors.New(fmt.Sprintf("`windows.[%v].name` is missing.", indexWindow))
		}
		for indexPane, pane := range window.Panes {
			if pane.Type == "" {
				return errors.New(fmt.Sprintf("`windows.[%v].panes.[%v].type` is missing.", indexWindow, indexPane))
			} else if pane.Type != Horizontal && pane.Type != Vertical {
				return errors.New(fmt.Sprintf("Invalid value for key `pane.type`.\nGot %v, expected `horizontal` or `vertical`.", pane.Type))
			}
		}
	}

	return nil
}

func (ws *Workspace) StartTmuxEnv(config *tmux.Config) error {
	tmux := tmux.Tmux{Config: config}

	if ws.Env != nil {
		tmux.Envs = &ws.Env
	}

	var sessionCwd string
	var err error
	if ws.CwdSession == "" {
		sessionCwd, err = os.Getwd()
		if err != nil {
			return err
		}
	} else {
		sessionCwd = ws.CwdSession
	}

	err = tmux.NewSession(ws.Session).SetCWD(sessionCwd).Execute()
	if err != nil {
		return err
	}

	wsWindow := ws.Windows[0]
	var initialWindowCwd string
	if wsWindow.Cwd == "" {
		initialWindowCwd = sessionCwd
	} else {
		initialWindowCwd = wsWindow.Cwd
	}

	// creating session actually create a window already
	err = tmux.RenameWindow(ws.Session, "0", wsWindow.Name).Execute()
	if err != nil {
		return err
	}

	err = tmux.SendKey(ws.Session, wsWindow.Name, "cd "+initialWindowCwd, "clear")
	if err != nil {
		return err
	}
	err = tmux.SendKey(ws.Session, wsWindow.Name, wsWindow.Commands...)
	if err != nil {
		return err
	}

	for _, wsPanes := range wsWindow.Panes {
		wsPanes.toTmux(&tmux, ws, &wsWindow, initialWindowCwd)
	}

	// initial window/session should run this too
	for _, wsWindow := range ws.Windows[1:] {
		wsWindow.toTmux(&tmux, ws, initialWindowCwd)
	}

	if ws.Follow {
		err := tmux.FollowSession(ws.Session + ":0").Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

func (wsWindow *WorkspaceWindow) toTmux(tmux *tmux.Tmux, ws *Workspace, highestCwd string) error {
	var windowCwd string
	if wsWindow.Cwd == "" {
		windowCwd = highestCwd
	} else {
		windowCwd = wsWindow.Cwd
	}

	if wsWindow.KeepAlive {
		wsWindow.Commands = append(wsWindow.Commands, "zsh")
	}

	err := tmux.NewWindow(ws.Session, wsWindow.Name).SetCWD(windowCwd).Execute(wsWindow.Commands...)
	if err != nil {
		return err
	}

	for _, wsPane := range wsWindow.Panes {
		err = wsPane.toTmux(tmux, ws, wsWindow, highestCwd)
		if err != nil {
			return err
		}
	}

	return nil
}
func (wsPane *WorkspacePane) toTmux(tmux *tmux.Tmux, ws *Workspace, wsWindow *WorkspaceWindow, highestCwd string) error {
	var paneCwd string
	if wsPane.Cwd == "" {
		paneCwd = highestCwd
	} else {
		paneCwd = wsPane.Cwd
	}

	if wsPane.KeepAlive {
		wsPane.Commands = append(wsPane.Commands, "zsh")
	}

	if wsPane.Type == Horizontal {
		err := tmux.NewSplitPaneHorizontal(ws.Session, wsWindow.Name).SetCWD(paneCwd).Execute(wsPane.Commands...)
		if err != nil {
			return err
		}
	} else {
		err := tmux.NewSplitPaneVertical(ws.Session, wsWindow.Name).SetCWD(paneCwd).Execute(wsPane.Commands...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ws *Workspace) StopTmuxEnv(config *tmux.Config) error {
	tmux := tmux.Tmux{Config: config}

	err := tmux.KillSession(ws.Session).Execute()
	if err != nil {
		return err
	}

	return nil
}

func newDefaultWorkspace(sessionName string) *Workspace {
	ws := Workspace{
		Session:    sessionName,
		Before:     []string{""},
		Stop:       []string{""},
		Env:        map[string]string{"VAR": "0"},
		Follow:     true,
		CwdSession: "./",
		Windows: []WorkspaceWindow{
			{
				Name:      "default_name",
				Cwd:       "~",
				Commands:  []string{"echo hello world"},
				KeepAlive: true,
				Panes: []WorkspacePane{
					{
						Type:      Horizontal,
						Cwd:       "./",
						KeepAlive: true,
						Commands:  []string{"echo hello world"},
					},
				},
			},
		},
	}

return &ws
}
