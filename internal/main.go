package main

import (
	"fmt"
	"github.com/YashKumarVerma/hentry-client/internal/config"
	"github.com/YashKumarVerma/hentry-client/internal/fs"
	"github.com/YashKumarVerma/hentry-client/internal/sensor"
	"github.com/YashKumarVerma/hentry-client/internal/watchman"
	"time"
)

func main() {
	// initialize all modules
	config.Init()
	sensor.Init()
	watchman.Init()

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

	if notFound == true {
		fmt.Println("Device not registered on Hentry, registration in process...")
		deviceDetails := device.CreateTeamScreen()
		teamNotFound, deviceAPIResponse := device.RegisterDeviceAPI(deviceDetails, UserTeam.ID)

		if teamNotFound == true {
			fmt.Println("Team does not exist")
		} else {
			fmt.Println("Device registered.")
			UserDevice = deviceAPIResponse
		}
	} else {
		fmt.Println("Device already registered.")
		UserDevice = deviceInfo
	}

	fmt.Println(UserTeam, UserDevice)

	success, folderNames := fs.ListFolders("./")
	if !success {
		fmt.Errorf("error reading folder names")
		return
	}

	// repeat the process based on config.frequency
	for i := 1; i >= 0; i++ {
		for _, folder := range folderNames {
			filesNotIgnored, _ := watchman.IndexAllFiles(folder)
			for _, i := range filesNotIgnored {
				_, fileDetails := fs.AnalyzeFile(i)
				watchman.ProcessFile(fileDetails)
			}
		}
		time.Sleep(1 * time.Second)
		fmt.Println("total entropy of project : ", watchman.AggregatorValue)
		watchman.ResetAggregator()
	}
}
