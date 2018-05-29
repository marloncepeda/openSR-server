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

// FormatSQLConnectionURL ...
func FormatSQLConnectionURL() string {

	return fmt.Sprintf("host=%s port=%v user=%s dbname=%s sslmode=%s password=%s",
		viper.Get("database.host"), viper.GetInt("database.port"),
		viper.Get("database.username"), viper.Get("database.database"),
		viper.Get("database.ssl"), viper.Get("database.password"))
}
