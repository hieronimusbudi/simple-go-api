package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	PORT       string `mapstructure:"PORT"`
	DBHOST     string `mapstructure:"DBHOST"`
	DBUSER     string `mapstructure:"DBUSER"`
	DBPASSWORD string `mapstructure:"DBPASSWORD"`
	DBNAME     string `mapstructure:"DBNAME"`
}

var c *Config

func SetConfig(configPath string) {
	var err error
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("config not found: %s", err.Error()))
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("could not parse config: %v", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Panicln("config file changed:", e.Name)
	})
	viper.WatchConfig()
}

func Get() *Config {
	return c
}
