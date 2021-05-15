package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Init seeds the configurations of the project.
func Init() {

	// initialize viper configurations
	viper.SetConfigName("hentry")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Custom configurations not found, using defaults")
	}

	// set the default as project deployment server
	viper.SetDefault("server.url", "http://40.119.165.213:8000")
	viper.SetDefault("feeder.url", "http://40.119.165.213:9000")

	// assing the configurations to the exported data member
	Load.Name = "hentry"
	Load.Server = viper.GetString("server.url")
	Load.Feeder = viper.GetString("feeder.url")
}

// Load exposes the configurations to other internal modules
var Load ConfigStruct
