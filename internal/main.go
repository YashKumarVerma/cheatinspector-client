package main

import (
	"fmt"
	"github.com/YashKumarVerma/hentry-client/internal/config"
	"github.com/YashKumarVerma/hentry-client/internal/team"
	"github.com/YashKumarVerma/hentry-client/internal/team/screens"
)

func main() {
	// initialize all modules
	config.Init()

	// show prompt to join a team or create a team
	teamOperation := screens.CreateOrJoinTeamScreen()
	if teamOperation == "create" {
		teamDetails := screens.CreateTeamScreen()
		teamAPIResponse := team.CreateTeamAPI(teamDetails)
		_ = teamAPIResponse
	} else {

	}

	fmt.Println("We're connected now.")
}
