package team

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	fmt.Println(string(body))
	if err != nil {
		log.Fatalln(err)
	}

	var parsedAPIResponse createTeamAPIResponse
	json.Unmarshal([]byte(body), &parsedAPIResponse)

	// print detailed log about API call
	fmt.Println("\nAPI Call Log ")
	fmt.Println("\nAPI Call Log : Target ", config.Load.Server+"/team")
	fmt.Println("\nAPI Call Log : Body : ", responseBody)
	fmt.Println("\nAPI Call Log : Params : ", team)
	fmt.Println("\nAPI Call Log : Response : Code ", parsedAPIResponse.Code)
	fmt.Println("\nAPI Call Log : Response : Error ", parsedAPIResponse.Err)
	fmt.Println("\nAPI Call Log : Response : Payload ", parsedAPIResponse.Payload.Devices)
	fmt.Println("\nAPI Call Log : Response : Payload ", parsedAPIResponse.Payload.ID)
	fmt.Println("\nAPI Call Log : Response : Payload ", parsedAPIResponse.Payload.Name)

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
	fmt.Println("API Response : ", parsedAPIResponse)
	json.Unmarshal([]byte(body), &parsedAPIResponse)

	return parsedAPIResponse.Payload
}
