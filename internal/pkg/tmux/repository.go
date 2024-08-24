package tmux

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Repository struct {
	Session string            `yaml:"session"`
	Before  []string          `yaml:"before,omitempty"`
	Stop    []string          `yaml:"stop,omitempty"`
	Env     map[string]string `yaml:"env,omitempty"`
	Windows []RepoWindow      `yaml:"windows"`
	Follow  bool              `yaml:"follow,omitempty"`
}

type RepoWindow struct {
	Name  string     `yaml:"name"`
	Cwd   string     `yaml:"cwd"`
	Panes []RepoPane `yaml:"panes"`
}

type RepoPane struct {
	Type     RepoPaneType `yaml:"type"`
	Cwd      string       `yaml:"cwd"`
	Commands []string     `yaml:"commands"`
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
		return repo, err
	}

	return repo, nil
}

func (repo *Repository) ParseToTmux(config *Config) error {
	tmux := Tmux{Config: config}

	if repo.Env != nil {
		tmux.Envs = &repo.Env
	}

	err := tmux.NewSession(repo.Session).Execute()
	if err != nil {
		return err
	}

	for index, repoWindow := range repo.Windows {
		if index == 0 {
			err := tmux.RenameWindow(repo.Session, "0", repoWindow.Name).Execute()
			if err != nil {
				return err
			}
		} else {
			err := tmux.NewWindow(repo.Session, repoWindow.Name).SetCWD(repoWindow.Cwd).Execute()
			if err != nil {
				return err
			}
		}

		for _, repoPane := range repoWindow.Panes {
			if repoPane.Type == Horizontal {
				err := tmux.NewSplitPaneHorizontal(repo.Session, repoWindow.Name).SetCWD(repoPane.Cwd).Execute(repoPane.Commands...)
				if err != nil {
					return err
				}
			} else {
				err := tmux.NewSplitPaneVertical(repo.Session, repoWindow.Name).SetCWD(repoPane.Cwd).Execute(repoPane.Commands...)
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
