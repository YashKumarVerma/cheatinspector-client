package config

// Load returns the current configurations
func Load() configStruct {
	data := configStruct{
		name:      "hentry",
		server:    "http://localhost:8000",
		frequency: 10,
	}

	return data
}
