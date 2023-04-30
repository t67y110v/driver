package config

import "github.com/spf13/viper"

type Config struct {
	ServerIP      string `mapstructure:"SERVER_IP"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	ServerCOMPort string `mapstructure:"SERVER_COM_PORT"`
	ScaleIP       string `mapstructure:"SCALE_IP"`
	ScalePort     string `mapstructure:"SCALE_PORT"`
	ScaleCOMPort  string `mapstructure:"SCALE_COM_PORT"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
