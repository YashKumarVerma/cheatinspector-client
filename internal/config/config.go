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
		fmt.Println("Configuration file named \"hentry.yaml\" not found, using defaults")
	} else {
		fmt.Println("Configurations loaded from file, overriding defaults.")
	}

	// knownConfigs is a list of files to ignore by default
	var knownConfigs = []string{".git", ".idea", ".cache", "node_modules", "dist"}
	var emptyList []string

	// set the default as project deployment server
	viper.SetDefault("app.server", "http://localhost:8000")
	viper.SetDefault("app.feeder", "http://localhost:9000")
	viper.SetDefault("app.debug", true)
	viper.SetDefault("ignore", emptyList)

	// join the pre-defined ignore list and the new list
	knownConfigs = append(knownConfigs, viper.GetStringSlice("ignore")...)

	fmt.Println("\nIgnoring Directories :")
	for _, val := range knownConfigs {
		fmt.Println("- " + val)
	}
	fmt.Println()

	// accessing the configurations to the exported data member
	Load.Name = "hentry"
	Load.Server = viper.GetString("app.server")
	Load.Feeder = viper.GetString("app.feeder")
	Load.Debug = viper.GetBool("app.debug")
	Load.Ignore = knownConfigs
}

// Load exposes the configurations to other internal modules
var Load ConfigStruct
