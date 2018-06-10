package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// LoadEnvironmentVariables ...
func LoadEnvironmentVariables() error {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		return errors.New("Error reading config file " + err.Error())
	}

	viper.Set("Connection", formatConnectionURL())

	return nil

}

func formatConnectionURL() string {

	host := viper.Get("database.host")
	port := viper.Get("database.port")
	user := viper.Get("database.username")
	db := viper.Get("database.database")
	pass := viper.Get("database.password")

	ssl := viper.Get("database.ssl")

	return fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s", host, port, user, db, pass, ssl)
}
