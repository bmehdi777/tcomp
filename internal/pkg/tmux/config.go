package tmux

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	TmuxSocketPath string `mapstructure:"tmux_socket_path"`
	ComposePath    string `mapstructure:"compose_path"`
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

	return nil
}
