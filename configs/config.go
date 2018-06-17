package configs

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Load ...
func Load() error {

	viper.SetConfigName("config")
	viper.AddConfigPath("../configs/")

	err := viper.ReadInConfig()

	if err != nil {
		return errors.New("Error reading config file " + err.Error())
	}

	viper.Set("Connection", db())

	return nil
}

func db() string {
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
