package server

type Response (type T) struct {
	code int
	error bool
	message string
	payload T
}