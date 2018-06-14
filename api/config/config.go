package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Init ...
func Init() error {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		return errors.New("Error reading config file " + err.Error())
	}

	viper.Set("Connection", url())

	return nil
}

func url() string {
	return fmt.Sprintf(
		"host=%s port=%v user=%s dbname=%s password=%s sslmode=%s",
		viper.Get("database.host"),
		viper.Get("database.port"),
		viper.Get("database.username"),
		viper.Get("database.database"),
		viper.Get("database.password"),
		viper.Get("database.ssl"),
	)
}
