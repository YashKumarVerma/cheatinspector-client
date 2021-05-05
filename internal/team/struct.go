package team

import "github.com/YashKumarVerma/hentry-client/internal/device"

// CreateTeamStruct to encapsulate all team data
type CreateTeamStruct struct {
	TeamID   string
	TeamName string
}

type JoinTeamStruct struct {
	TeamID string
}

// Team contains all details of given team
type Team struct {
	Name    string          `json:"friendlyName"`
	ID      string          `json:"id"`
	Devices []device.Device `json:"devices"`
}

type createTeamAPIResponse struct {
	Err     bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload Team   `json:"payload"`
}

type getTeamDetailsAPIResponse struct {
	Err     bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload Team   `json:"payload"`
}
