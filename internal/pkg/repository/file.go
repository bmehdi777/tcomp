package repository

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bmehdi777/tmuxcompose/internal/pkg/tmux"
)

func ListRepository(config *tmux.Config) ([]string, error) {
	files, err := os.ReadDir(config.ComposePath)
	if err != nil {
		return nil, err
	}

	var filesNames []string
	for _, file := range files {
		if !file.IsDir() {
			filesNames = append(filesNames, file.Name())
		}
	}

	return filesNames, nil
}

func GetFileRepoPath(name string, config *tmux.Config) (string, error) {
	files, err := os.ReadDir(config.ComposePath)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			if strings.Contains(filename, name) {
				return fmt.Sprintf("%v/%v", config.ComposePath, filename), nil
			}
		}
	}

	return "", errors.New(fmt.Sprintf("No repository file named `%v` has been found.", name))
}