package form

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// CreateOrJoinTeam shows UI to select user choice on start-up
func CreateOrJoinTeam() string {

	// declare choices for interface
	prompt := promptui.Select{
		Label: "Team Status",
		Items: []string{"Create a team", "Join a team"},
	}

	// extract data from selection
	_, result, err := prompt.Run()

	// if error in selection, show error and move to team creation
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "create"
	}

	if result == "Create a team" {
		return "create"
	}

	return "join"
}
