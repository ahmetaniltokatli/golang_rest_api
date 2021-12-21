package models

type Response struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
	Data         string `json:"data"`
}
