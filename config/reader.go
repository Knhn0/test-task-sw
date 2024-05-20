package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Read() Config {
	env, configName := readEnv()

	viper.SetConfigName(configName)
	viper.SetConfigType("json")
	viper.AddConfigPath("./.config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var c Config
	c.Env = env
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return c
}

const serviceEnvVarName = "ENV"

type Env string

const (
	EnvProd  = Env("prod")
	EnvTest  = Env("test")
	EnvLocal = Env("local")
)

func readEnv() (Env, string) {
	env := Env(os.Getenv(serviceEnvVarName))

	var configName string
	switch env {
	case EnvProd:
		configName = "prod"
	case EnvTest:
		configName = "test"
	default:
		env = EnvLocal
		configName = "local"
	}

	return env, configName
}
