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
	Windows []RepoWindow        `yaml:"windows"`
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
