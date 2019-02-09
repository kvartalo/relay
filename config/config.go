package config

import (
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

type Config struct {
	Keystorage struct {
		Address  string
		Password string
	}
	Server struct {
		Port string
	}
	Web3 struct {
		Url string
	}
	Contracts struct {
		Token string
	}
}

var C Config

func MustRead(c *cli.Context) error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	err := viper.Unmarshal(&C)
	if err != nil {
		return err
	}
	return nil
}
