package config

import "flag"

// Load returns the current configurations
func Init() {

	serverURL := flag.String("api", "http://40.119.165.213:8000", "URL of main API server")
	feeder := flag.String("feed", "http://40.119.165.213:9000", "URL of main API server")

	data := ConfigStruct{
		Name:      "hentry",
		Server:    *serverURL,
		Feeder:    *feeder,
		Frequency: 10,
		FileName:  ".hentryrc",
	}

	Load.Name = data.Name
	Load.Server = data.Feeder
	Load.Frequency = data.Frequency
	Load.Feeder = data.Feeder
}

var Load ConfigStruct
