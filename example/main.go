package main

import (
	"example/wiring"
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
	// timescaleDB := config.InitTimeSeriesDatabase()

	mux := http.NewServeMux()
	apiPrinter := helper.NewApiPrinter("example Swagger Doc API", "example swagger open API documentation genarated automaticaly by system")

	ssed := helper.NewSSEFDefault()
	mux.HandleFunc("GET /example/sse", ssed.HandleSSE)

	cronjobDB, err := cronjob.NewMariaDBStorage(mariaDB)
	if err != nil {
		panic(err.Error())
	}

	cj := cronjob.NewCronJob(nil, cronjobDB)

	wiring.SetupDependency(mariaDB, mux, jwtToken, apiPrinter, cj, ssed)

	apiPrinter.PrintAPIDataTable().PublishAPI(mux, os.Getenv("SERVER_URL"), "/openapi")

	helper.StartServerWithGracefullyShutdown(mux)
}
