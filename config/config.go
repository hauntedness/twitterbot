package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	BotConf BotConfig
}

func GetTwitterConfig() *config {
	path := os.Getenv("TWITTER_CONFIG_PATH")
	if path == "" {
		dir, _ := os.UserConfigDir()
		path = dir + `\twitter\.config\twitter.toml`
	}
	conf := config{}
	_, err := toml.DecodeFile(path, &conf)
	if err != nil {
		return nil
	}
	return &conf
}

type BotConfig struct {
	Token string
}

func GetBotConfig() BotConfig {
	c := GetTwitterConfig()
	return c.BotConf
}
