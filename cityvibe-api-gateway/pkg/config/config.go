package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT          string `mapstructure:"PORT"`
	AdminSvcUrl   string `mapstructure:"AdminSvcUrl"`
	OrderSvcUrl   string `mapstructure:"OrderSvcUrl"`
	PaymentSvcUrl string `mapstructure:"PaymentSvcUrl"`
	ProductSvcUrl string `mapstructure:"ProductSvcUrl"`
	UserSvcUrl    string `mapstructure:"UserSvcUrl"`
	CartSvcUrl    string `mapstructurea:"CartSvcUrl"`
	KEY           string `mapstructure:"KEY"`
}

func NewConfig() (config *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return
}
