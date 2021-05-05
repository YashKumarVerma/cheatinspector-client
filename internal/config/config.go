package config

// Load returns the current configurations
func Init() {
	data := ConfigStruct{
		Name:      "hentry",
		Server:    "http://localhost:8000",
		Frequency: 10,
		FileName:  ".hentryrc",
	}

	Load.Name = data.Name
	Load.Server = data.Server
	Load.Frequency = data.Frequency
}

var Load ConfigStruct
