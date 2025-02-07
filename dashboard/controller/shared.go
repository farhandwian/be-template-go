package controller

import (
	"net/http"
	"shared/helper"
)

type Controller struct {
	Mux *http.ServeMux
	JWT helper.JWTTokenizer
	// Cfg helper.AppConfig
}
