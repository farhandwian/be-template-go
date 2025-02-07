package main

import (
	"monitoring/wiring"
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

	// db
	mariaDB := config.InitMariaDatabase()

	mux := http.NewServeMux()
	apiPrinter := helper.NewApiPrinter()

	wiring.SetupDependency(mariaDB, mux, apiPrinter)

	apiPrinter.PrintAPIDataTable().PublishAPI(mux, os.Getenv("SERVER_URL"), "/openapi")

	helper.StartServerWithGracefullyShutdown(mux)
}
