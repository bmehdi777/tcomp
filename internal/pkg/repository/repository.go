package repository

import (
	"errors"
	"fmt"
	"os"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
	"gopkg.in/yaml.v3"
)

type Repository struct {
	Session    string            `yaml:"session"`          // mandatory
	Before     []string          `yaml:"before,omitempty"` // optional
	Stop       []string          `yaml:"stop,omitempty"`   // optional
	Env        map[string]string `yaml:"env,omitempty"`    // optional
	Follow     bool              `yaml:"follow"`           // optional
	CwdSession string            `yaml:"cwd"`              // optional
	Windows    []RepoWindow      `yaml:"windows"`          // mandatory
}

type RepoWindow struct {
	Name      string     `yaml:"name"`               // mandatory
	Cwd       string     `yaml:"cwd"`                // optional - get current working dir
	Commands  []string   `yaml:"commands,omitempty"` // optional - do nothing
	Panes     []RepoPane `yaml:"panes,omitempty"`    // optional
	KeepAlive bool       `yaml:"keep_alive"`         //optional
}

type RepoPane struct {
	Type      RepoPaneType `yaml:"type,omitempty"`     // mandatory
	Cwd       string       `yaml:"cwd"`                // optional
	Commands  []string     `yaml:"commands,omitempty"` // optional
	KeepAlive bool         `yaml:"keep_alive"`         //optional
}

type RepoPaneType string

const (
	Horizontal RepoPaneType = "horizontal"
	Vertical   RepoPaneType = "vertical"
)

func ReadRepository(pathfile string) (Repository, error) {
	repo := Repository{}
	data, err := os.ReadFile(pathfile)
	if err != nil {
		return repo, err
	}

	err = yaml.Unmarshal(data, &repo)
	if err != nil {
		fmt.Println("Err :%w")
		return repo, err
	}

	err = repo.verifyRepository()
	if err != nil {
		return Repository{}, err
	}

	return repo, nil
}

func (repo *Repository) verifyRepository() error {
	if repo.Session == "" {
		return errors.New(fmt.Sprintf("`session` is missing."))
	}
	if repo.Windows == nil {
		return errors.New(fmt.Sprintf("`windows` is missing."))
	}

	for indexWindow, window := range repo.Windows {
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

func (repo *Repository) StartTmuxEnv(config *tmux.Config) error {
	tmux := tmux.Tmux{Config: config}

	if repo.Env != nil {
		tmux.Envs = &repo.Env
	}

	var sessionCwd string
	var err error
	if repo.CwdSession == "" {
		sessionCwd, err = os.Getwd()
		if err != nil {
			return err
		}
	} else {
		sessionCwd = repo.CwdSession
	}

	err = tmux.NewSession(repo.Session).SetCWD(sessionCwd).Execute()
	if err != nil {
		return err
	}

	for index, repoWindow := range repo.Windows {
		var windowCwd string
		if repoWindow.Cwd == "" {
			windowCwd = sessionCwd
		} else {
			windowCwd = repoWindow.Cwd
		}

		if repoWindow.KeepAlive {
			// todo: change to current shell used
			repoWindow.Commands = append(repoWindow.Commands, "zsh")
		}

		// window already exist when creating session
		if index == 0 {
			err := tmux.RenameWindow(repo.Session, "0", repoWindow.Name).Execute()
			if err != nil {
				return err
			}

			// window is already created so if I want to move or execute anything...
			// todo: find a better solution
			err = tmux.SendKey(repo.Session, repoWindow.Name, "cd "+windowCwd, "clear")
			if err != nil {
				return err
			}
			err = tmux.SendKey(repo.Session, repoWindow.Name, repoWindow.Commands...)
			if err != nil {
				return err
			}
		} else {
			err := tmux.NewWindow(repo.Session, repoWindow.Name).SetCWD(windowCwd).Execute(repoWindow.Commands...)
			if err != nil {
				return err
			}
		}

		for _, repoPane := range repoWindow.Panes {
			var paneCwd string
			if repoPane.Cwd == "" {
				paneCwd = windowCwd
			} else {
				paneCwd = repoPane.Cwd
			}

			if repoPane.KeepAlive {
				repoPane.Commands = append(repoPane.Commands, "zsh")
			}

			if repoPane.Type == Horizontal {
				err := tmux.NewSplitPaneHorizontal(repo.Session, repoWindow.Name).SetCWD(paneCwd).Execute(repoPane.Commands...)
				if err != nil {
					return err
				}
			} else {
				err := tmux.NewSplitPaneVertical(repo.Session, repoWindow.Name).SetCWD(paneCwd).Execute(repoPane.Commands...)
				if err != nil {
					return err
				}
			}

		}
	}

	if repo.Follow {
		err := tmux.FollowSession(repo.Session + ":0").Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) StopTmuxEnv(config *tmux.Config) error {
	tmux := tmux.Tmux{Config: config}

	err := tmux.KillSession(repo.Session).Execute()
	if err != nil {
		return err
	}

	return nil
}
