package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MySQLdsn                     string
	FirebaseServieAccountKeyPath string
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := &Config{
		MySQLdsn:                     viper.GetString("mysql.dsn"),
		FirebaseServieAccountKeyPath: viper.GetString("firebase.service_account_key_path"),
	}
	fmt.Printf("config: %v \n", config)
	return config, nil
}
