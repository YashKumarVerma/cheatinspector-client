package config

// Load returns the current configurations
func Init() {
	data := ConfigStruct{
		Name:      "hentry",
		Server:    "http://localhost:8000",
		Feeder:    "http://localhost:9898",
		Frequency: 10,
		FileName:  ".hentryrc",
	}

	Load.Name = data.Name
	Load.Server = data.Server
	Load.Frequency = data.Frequency
	Load.Feeder = data.Feeder
}

var Load ConfigStruct
