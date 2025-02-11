package main

import (
	"context"
	"encoding/json"
	example "example/wiring"
	iam "iam/wiring"

	"net/http"
	"os"
	"shared/config"
	"shared/helper"
	"shared/helper/cronjob"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	jwtToken, err := helper.NewJWTTokenizer(os.Getenv("TOKEN"))
	if err != nil {
		panic(err.Error())
	}

	ec := iam.InitEmail()

	mangantiDB := config.InitMariaDatabase()

	// dashboard.InitSeeder(mangantiDB)

	// tsdb := config.InitTimeSeriesDatabase()

	iam.CreateAdminIfNotExists(mangantiDB)

	mux := http.NewServeMux()

	// sseBigboard := helper.NewSSEFDefault()
	// mux.HandleFunc("GET /bigboard/sse", sseBigboard.HandleSSE)

	sseDashboard := helper.NewSSEFDefault()
	mux.HandleFunc("GET /dashboard/sse", sseDashboard.HandleSSE)

	apiPrinter := helper.NewApiPrinter("", "")

	cronjobDB, err := cronjob.NewMariaDBStorage(mangantiDB)
	if err != nil {
		panic(err.Error())
	}

	cj := cronjob.NewCronJob(nil, cronjobDB)

	iam.SetupDependencyWithDatabaseAndEmail(apiPrinter, mux, jwtToken, mangantiDB, ec)
	example.SetupDependency(mangantiDB, mux, jwtToken, apiPrinter, cj, sseDashboard)
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})

	apiPrinter.PrintAPIDataTable().PublishAPI(mux, os.Getenv("SERVER_URL"), "/openapi")

	if err := cj.Start(context.Background()); err != nil {
		panic(err.Error())
	}

	helper.StartServerWithGracefullyShutdown(mux)

}
