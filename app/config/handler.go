package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// LoadConfigurationVariables function return a yaml data
func LoadConfigurationVariables() {

	viper.SetConfigName("config")
	//viper.AddConfigPath("../")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
}

// FormatSQLAddressConnection ...
func FormatSQLAddressConnection() string {
	return fmt.Sprintf("%s:%v", viper.Get("database.host"), viper.Get("database.port"))
}

// GetDatabaseUsername ...
func GetDatabaseUsername() string { return fmt.Sprintf("%s", viper.Get("database.username")) }

// GetDatabasePassword ...
func GetDatabasePassword() string { return fmt.Sprintf("%s", viper.Get("database.password")) }

// GetDatabaseDatabase ...
func GetDatabaseDatabase() string { return fmt.Sprintf("%s", viper.Get("database.database")) }
