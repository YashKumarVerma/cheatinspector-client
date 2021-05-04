package main

import (
	"fmt"
	"github.com/YashKumarVerma/hentry-client/internal/config"
	"github.com/YashKumarVerma/hentry-client/internal/sensor"
	"github.com/YashKumarVerma/hentry-client/internal/team"
)

func main() {
	// initialize all modules
	config.Init()
	sensor.Init()

	// show prompt to join a team or create a team
	teamOperation := team.CreateOrJoinTeamScreen()
	if teamOperation == "create" {
		teamDetails := team.CreateTeamScreen()
		teamAPIResponse := team.CreateTeamAPI(teamDetails)
		_ = teamAPIResponse
	} else {
		teamDetails := team.JoinTeamScreen()
		teamAPIResponse := team.GetTeamDetailAPI(teamDetails.TeamID)
		_ = teamAPIResponse
	}

	// check if current device is registered on the server or not
	userDeviceInfo := sensor.Load
	fmt.Println("Logged in on device : " + userDeviceInfo.MachineID)

	fmt.Println("We're connected now.")
}
