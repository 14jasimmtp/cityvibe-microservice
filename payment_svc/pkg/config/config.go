package config

import "github.com/spf13/viper"

type Config struct {
	DB_URL             string `mapstructure:"DB_URL"`
	PORT               string `mapstructure:"PORT"`
	KEY_ID_FOR_PAY     string `mapstructure:"KEY_ID_FOR_PAY"`
	SECRET_KEY_FOR_PAY string `mapstructure:"SECRET_KEY_FOR_PAY"`
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
