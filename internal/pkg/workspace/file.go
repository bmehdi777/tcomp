package workspace

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/bmehdi777/tcomp/internal/pkg/tmux"
	"gopkg.in/yaml.v3"
)

func ListWorkspace(config *tmux.Config) ([]string, error) {
	files, err := os.ReadDir(config.ComposePath)
	if err != nil {
		return nil, err
	}

	var filesNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			filesNames = append(filesNames, strings.TrimSuffix(fileName, filepath.Ext(fileName)))
		}
	}

	return filesNames, nil
}

func GetFileWorkspacePath(name string, config *tmux.Config) (string, error) {
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

	return "", errors.New(fmt.Sprintf("No workspace file named `%v` has been found.", name))
}

func CreateNewWorkspaceFile(name string, config *tmux.Config) (string, error) {
	fullPath := path.Join(config.ComposePath, name)
	fullPath = fullPath + ".yml"
	_, err := os.Stat(fullPath)

	if err == nil {
		if errors.Is(err, os.ErrExist) {
			return "", errors.New("File already exist.")
		}
		return "", err
	}

	newFile, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	defaultWs := newDefaultWorkspace(name)
	buf, err := yaml.Marshal(&defaultWs)
	if err != nil {
		return "", err
	}

	_, err = newFile.Write(buf)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}

func OpenWorkspaceFileWithEditor(filepath string, config *tmux.Config) error {
	editor := os.Getenv("EDITOR")

	cmd := exec.Command(editor, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
