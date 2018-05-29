package app

import (
	"github.com/ctreminiom/scientific-logs-api/config"
	"github.com/spf13/viper"
)

// Config function return a yaml data
func Config() (data *config.Configuration, err error) {

	viper.SetConfigName("config")
	//viper.AddConfigPath("../")
	viper.AddConfigPath(".")

	var env config.Configuration
	var validator = viper.ReadInConfig()

	if validator != nil {
		return nil, validator
	}

	yaml := viper.Unmarshal(&env)

	if yaml != nil {
		return nil, yaml
	}

	return &env, nil
}

/*
// FormatSQLConnectionURL ...
func FormatSQLConnectionURL(data *config.Configuration) string {

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		data.Database.Host, data.Database.Port,
		data.Database.Username, data.Database.Database,
		data.Database.SSL, data.Database.Password)
}


*/
