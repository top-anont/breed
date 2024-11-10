package config

import "github.com/spf13/viper"

var AppConfig *ConfigStructure

type ConfigStructure struct {
	Mode             string `mapstructure:"MODE"`
	Port             string `mapstructure:"PORT"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
}

func LoadConfig(path string) (config *ConfigStructure, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	errViperReadConfig := viper.ReadInConfig()
	if errViperReadConfig != nil {
		return
	}

	errViperUnmarshal := viper.Unmarshal(&config)
	if errViperUnmarshal != nil {
		return
	}

	return
}
