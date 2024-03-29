package config

import "github.com/spf13/viper"

type Config struct {
	DB_URL string `mapstructure:"DB_URL"`
	PORT   string `mapstructure:"PORT"`
	KEY    string `mapstructure:"KEY"`
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
