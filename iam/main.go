package main

import (
	"fmt"
	"iam/wiring"
	"net/http"
	"os"
	"shared/config"
	"shared/helper"

	"github.com/rs/cors"
)

func main() {

	jwtToken, err := helper.NewJWTTokenizer(os.Getenv("TOKEN"))
	if err != nil {
		panic(err.Error())
	}

	db := config.InitMariaDatabase()

	wiring.CreateAdminIfNotExists(db)

	mux := http.NewServeMux()

	apiPrinter := helper.NewApiPrinter()

	wiring.SetupDependencyWithDatabase(apiPrinter, mux, jwtToken, db)

	apiPrinter.PrintAPIDataTable().PublishAPI(mux, os.Getenv("SERVER_URL"), "/openapi")

	panic(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors.AllowAll().Handler(mux)))

}
