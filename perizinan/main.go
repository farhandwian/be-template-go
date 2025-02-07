package main

import (

	// "perizinan/wiring"
	// "shared/helper"

	"encoding/json"
	"net/http"
	"os"
	"perizinan/wiring"
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
	mariaDB := wiring.InitMariaDatabase()

	mux := http.NewServeMux()
	apiPrinter := helper.NewApiPrinter()

	jwtToken, err := helper.NewJWTTokenizer(os.Getenv("TOKEN"))
	if err != nil {
		panic(err.Error())
	}

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	})

	wiring.SetupDependency(mariaDB, mux, apiPrinter, jwtToken)

	apiPrinter.PrintAPIDataTable().PublishAPI(mux, os.Getenv("SERVER_URL"), "/openapi")

	helper.StartServerWithGracefullyShutdown(mux)
}
