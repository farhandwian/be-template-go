package main

import (
	"dashboard/wiring"
	"net/http"
	"os"
	"shared/config"
	"shared/helper"
	"shared/helper/cronjob"

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

	ssed := helper.NewSSEFDefault()
	mux.HandleFunc("GET /dashboard/sse", ssed.HandleSSE)

	sseb := helper.NewSSEFDefault()
	mux.HandleFunc("GET /bigboard/sse", sseb.HandleSSE)

	cronjobDB, err := cronjob.NewMariaDBStorage(mariaDB)
	if err != nil {
		panic(err.Error())
	}

	cj := cronjob.NewCronJob(nil, cronjobDB)

	wiring.SetupDependency(mariaDB, timescaleDB, mux, jwtToken, apiPrinter, cj, sseb, ssed)

	apiPrinter.PrintAPIDataTable().PublishAPI(mux, os.Getenv("SERVER_URL"), "/openapi")

	helper.StartServerWithGracefullyShutdown(mux)
}
