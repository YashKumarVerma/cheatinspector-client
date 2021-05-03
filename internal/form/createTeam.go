package form

import (
	"errors"
	"fmt"
	"os/user"

	"github.com/manifoldco/promptui"
)

//  to get the team ID input from user
func getTeamId() string {
	validate := func(input string) error {
		if len(input) < 4 {
			return errors.New("team id should be at-least 4 characters long")
		}
		return nil
	}

	var teamID string
	u, err := user.Current()
	if err == nil {
		teamID = u.Username
	}

	prompt := promptui.Prompt{
		Label:    "Team ID",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return teamID
	}

	return result
}

//  to get the team name input from user
func getTeamName() string {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("friendly team name should be at-least 6 characters")
		}
		return nil
	}

	var teamID string
	u, err := user.Current()
	if err == nil {
		teamID = u.Username
	}

	prompt := promptui.Prompt{
		Label:    "Team Name",
		Validate: validate,
		Default:  "team " + teamID,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return teamID
	}

	return result
}

// CreateOrJoinTeam shows UI to select user choice on start-up
func CreateTeam() CreateTeamStruct {

	teamDetails := CreateTeamStruct{
		TeamID:   getTeamId(),
		TeamName: getTeamName(),
	}

	fmt.Println(teamDetails.TeamID)
	fmt.Println(teamDetails.TeamName)

	return teamDetails
}
