package main

import (
	"fmt"
	"github.com/YashKumarVerma/hentry-client/internal/config"
	"github.com/YashKumarVerma/hentry-client/internal/device"
	"github.com/YashKumarVerma/hentry-client/internal/sensor"
	"github.com/YashKumarVerma/hentry-client/internal/team"
)

func main() {
	// initialize all modules
	config.Init()
	sensor.Init()

	// holders for user data
	var UserTeam team.Team
	var UserDevice device.Device

	// register or join a team
	teamOperation := team.CreateOrJoinTeamScreen()
	if teamOperation == "create" {
		teamDetails := team.CreateTeamScreen()
		teamAPIResponse := team.CreateTeamAPI(teamDetails)
		UserTeam = teamAPIResponse
	} else {
		teamDetails := team.JoinTeamScreen()
		teamAPIResponse := team.GetTeamDetailAPI(teamDetails.TeamID)
		UserTeam = teamAPIResponse
	}

	// load all sensor data
	userDeviceInfo := sensor.Load
	notFound, deviceInfo := device.GetDeviceDetailAPI(userDeviceInfo.MachineID)

	if notFound {
		fmt.Println("Device not registered on Hentry, registration in process...")
		deviceDetails := device.CreateTeamScreen()
		deviceAPIResponse := device.RegisterDeviceAPI(deviceDetails, UserTeam.ID)
		fmt.Println("Device registered.")
		UserDevice = deviceAPIResponse
	} else {
		fmt.Println("Device alreadt registered.")
		UserDevice = deviceInfo
	}

	fmt.Println(UserTeam, UserDevice)
}
