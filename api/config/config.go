package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Load ...
func Load() error {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		return errors.New("Error reading config file " + err.Error())
	}

	return nil

}

// Parameters ...
type Parameters struct{ Username, Password, Addr, Database string }

// Fetch ...
func Fetch() *Parameters {

	return &Parameters{
		Username: username(),
		Password: password(),
		Addr:     address(),
		Database: database(),
	}

}

func username() string { return fmt.Sprintf("%s", viper.Get("database.username")) }
func password() string { return fmt.Sprintf("%s", viper.Get("database.password")) }
func database() string { return fmt.Sprintf("%s", viper.Get("database.database")) }

func address() string {
	return fmt.Sprintf("%s:%v", viper.Get("database.host"), viper.Get("database.port"))
}
