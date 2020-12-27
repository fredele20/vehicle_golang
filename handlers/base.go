package handlers

import "vehicle_golang/models"

const (
	BaseRoute = "/api"
)

type errorRes struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statuscode"`
	Error_Des  string `json:"error_description"`
}

type basicResponse struct {
	Message string `json:"message"`
}

type loginRes struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

type signupReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
