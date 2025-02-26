package main

import (
	"fmt"
	"iam/wiring"
	"net/http"
	"os"
	"shared/config"
	"shared/helper"
	oryHelper "shared/helper/ory"
	ketoHelper "shared/helper/ory/keto"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	db := config.InitMariaDatabase()

	wiring.CreateAdminIfNotExists(db)

	mux := http.NewServeMux()

	ketoClient := ketoHelper.SetupKetoGRPCClient()

	oryClient := oryHelper.SetupOryClient()

	apiPrinter := helper.NewApiPrinter("I AM", "I AM api documentation")

	wiring.SetupDependencyWithDatabase(apiPrinter, mux, ketoClient, oryClient, jwtToken, db)

	apiPrinter.PrintAPIDataTable().PublishAPI(mux, os.Getenv("SERVER_URL"), "/openapi")

	panic(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors.AllowAll().Handler(mux)))

}
