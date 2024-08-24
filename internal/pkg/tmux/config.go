package tmux

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	TmuxSocketPath string `mapstructure:"tmux_socket_path"`
	ComposePath    string `mapstructure:"compose_repository"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/tcomp/")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := createDefaultConfig()
			if err != nil {
				fmt.Println("Error : ", err)
				return nil, err
			}
		} else {
			fmt.Println("Error : ", err)
			return nil, err
		}
	}

	var config Config
	viper.Unmarshal(&config)

	return &config, nil
}

func createDefaultConfig() error {
	fmt.Println("Create initial config")
	configPath := filepath.Join(os.Getenv("HOME"), ".config/tcomp")
	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}

	repoPath := filepath.Join(configPath, "repository/")
	err = os.MkdirAll(repoPath, os.ModePerm)
	if err != nil {
		return err
	}

	viper.Set("tmux_socket_path", "/tmp/tmux-1000/default")
	viper.Set("compose_repository", filepath.Join(configPath, "/repository"))
	fmt.Println("test")
	viper.SafeWriteConfig()
	return nil
}
