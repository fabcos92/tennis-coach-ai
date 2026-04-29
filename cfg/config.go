package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Env  string `mapstructure:"env"`
		Name string `mapstructure:"name"`
	}

	HTTP struct {
		Host string `mapstructure:"host"`
		Port uint64 `mapstructure:"port"`
	} `mapstructure:"http"`
}

func Load(cfgFile string) *Config {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
	} else {
		log.Fatalf("failed to read config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return &cfg
}
