package team

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/YashKumarVerma/hentry-client/internal/config"
)

// CreateTeamAPI makes call to server to create new team
func CreateTeamAPI(team CreateTeamStruct) Team {

	// structure the data that needs to be sent
	postBody, _ := json.Marshal(map[string]string{
		"id":           team.TeamID,
		"friendlyName": team.TeamName,
	})
	responseBody := bytes.NewBuffer(postBody)

	// make the request
	resp, err := http.Post(config.Load.Server+"/team", "application/json", responseBody)

	// check if non 200 response
	if err != nil {
		log.Fatalf("An Error Encountered %v", err)
	}
	defer resp.Body.Close()

	// read response body as required
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var parsedAPIResponse createTeamAPIResponse
	//sb := string(body)
	json.Unmarshal([]byte(body), &parsedAPIResponse)

	return parsedAPIResponse.Payload
}

func GetTeamDetailAPI(teamID string) Team {
	resp, err := http.Get(config.Load.Server + "/team/" + teamID)

	// check if non 200 response
	if err != nil {
		log.Fatalf("An Error Encountered %v", err)
	}
	defer resp.Body.Close()

	// read response body as required
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var parsedAPIResponse getTeamDetailsAPIResponse
	json.Unmarshal([]byte(body), &parsedAPIResponse)

	return parsedAPIResponse.Payload
}
