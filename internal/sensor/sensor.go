package sensor

// Init stars the sensor calls to collect data
func Init() {
	Load.MachineID = getMachineId()
}

// Load returns current sensor data for other modules
var Load Sensor
