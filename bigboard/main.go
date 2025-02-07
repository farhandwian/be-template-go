package main

import (
	"bigboard/wiring"
	"net/http"
	"os"
	"shared/config"
	"shared/helper"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(".env file not found")
	}

}

func main() {

	jwtToken, err := helper.NewJWTTokenizer(os.Getenv("TOKEN"))
	if err != nil {
		panic(err.Error())
	}

	// db
	mariaDB := config.InitMariaDatabase()
	timescaleDB := config.InitTimeSeriesDatabase()

	mux := http.NewServeMux()
	apiPrinter := helper.NewApiPrinter()

	sse := helper.NewSSEFDefault()
	mux.HandleFunc("GET /sse", sse.HandleSSE)

	wiring.SetupDependency(mariaDB, timescaleDB, mux, jwtToken, apiPrinter, sse)

	apiPrinter.PrintAPIDataTable().PublishAPI(mux, os.Getenv("SERVER_URL"), "/openapi")

	helper.StartServerWithGracefullyShutdown(mux)
}
