package main

import (
	bigboard "bigboard/wiring"
	"context"
	dashboard "dashboard/wiring"
	"encoding/json"
	iam "iam/wiring"
	monitoring "monitoring/wiring"

	"net/http"
	"os"
	perizinan "perizinan/wiring"
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

	tsdb := config.InitTimeSeriesDatabase()

	iam.CreateAdminIfNotExists(mangantiDB)

	mux := http.NewServeMux()

	sseBigboard := helper.NewSSEFDefault()
	mux.HandleFunc("GET /bigboard/sse", sseBigboard.HandleSSE)

	sseDashboard := helper.NewSSEFDefault()
	mux.HandleFunc("GET /dashboard/sse", sseDashboard.HandleSSE)

	apiPrinter := helper.NewApiPrinter()

	cronjobDB, err := cronjob.NewMariaDBStorage(mangantiDB)
	if err != nil {
		panic(err.Error())
	}

	cj := cronjob.NewCronJob(nil, cronjobDB)

	iam.SetupDependencyWithDatabaseAndEmail(apiPrinter, mux, jwtToken, mangantiDB, ec)
	bigboard.SetupDependency(mangantiDB, tsdb, mux, jwtToken, apiPrinter, sseBigboard)
	dashboard.SetupDependency(mangantiDB, tsdb, mux, jwtToken, apiPrinter, cj, sseBigboard, sseDashboard)
	monitoring.SetupDependency(mangantiDB, mux, apiPrinter)
	perizinan.SetupDependency(mangantiDB, mux, apiPrinter, jwtToken)

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
