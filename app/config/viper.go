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

// Database ...
func Database() map[string]string {

	values := make(map[string]string)

	values["addr"] = address()
	values["username"] = username()
	values["password"] = password()
	values["database"] = database()

	return values

}

func address() string {
	return fmt.Sprintf("%s:%v", viper.Get("database.host"), viper.Get("database.port"))
}
func username() string { return fmt.Sprintf("%s", viper.Get("database.username")) }
func password() string { return fmt.Sprintf("%s", viper.Get("database.password")) }
func database() string { return fmt.Sprintf("%s", viper.Get("database.database")) }
