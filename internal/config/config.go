package config

import (
	"fmt"
)

// Load returns the current configurations
func Init() {

	serverURL := "http://40.119.165.213:8000"
	feeder := "http://40.119.165.213:9000"
	// feeder := "http://localhost:9898"

	data := ConfigStruct{
		Name:      "hentry",
		Server:    serverURL,
		Feeder:    feeder,
		Frequency: 10,
		FileName:  ".hentryrc",
	}

	fmt.Println("Current Configs: ")
	fmt.Println("Name : ", data.Name)
	fmt.Println("Frequency : ", data.Frequency)
	fmt.Println("Server : ", data.Server)
	fmt.Println("Feeder : ", data.Feeder)
	fmt.Println("\n\n")

	Load.Name = data.Name
	Load.Server = data.Server
	Load.Frequency = data.Frequency
	Load.Feeder = data.Feeder
}

var Load ConfigStruct
