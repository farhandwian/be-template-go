package flowtest

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"iam/model"
// 	"iam/wiring"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"shared/config"
// 	"shared/helper"

// 	"testing"

// 	"github.com/joho/godotenv"
// )

// func Test01(t *testing.T) {
// 	if err := godotenv.Load(); err != nil {
// 		panic(".env file not found")
// 	}
// 	jwtToken, _ := helper.NewJWTTokenizer(os.Getenv("TOKEN"))

// 	db := config.InitMariaDatabase()

// 	resetAllDatabaseForTestingPurpose(db)

// 	wiring.CreateAdminIfNotExists(db)

// 	mux := http.NewServeMux()

// 	apiPrinter := helper.NewApiPrinter("", "")

// 	wiring.SetupDependencyWithDatabase(apiPrinter, mux, jwtToken, db)

// 	// check user
// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "admin@mail.com")
// 		printUser("before login", user, "admin1234", "1234", "")
// 	}

// 	// login
// 	{
// 		body := struct {
// 			Email    model.Email `json:"email"`
// 			Password string      `json:"password"`
// 		}{
// 			Email:    model.Email("admin@mail.com"),
// 			Password: "admin1234",
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> login finish with response code %d\n", rr.Code)
// 	}

// 	// check user
// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "admin@mail.com")
// 		printUser("after login ", user, "admin1234", "1234", "123456")
// 	}

// 	accessToken := ""
// 	refreshToken := ""

// 	// login otp
// 	{
// 		body := struct {
// 			Email model.Email `json:"email"`
// 			OTP   string      `json:"otp"`
// 		}{
// 			Email: model.Email("admin@mail.com"),
// 			OTP:   "123456",
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/auth/login/otp", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		type Response struct {
// 			RefreshToken string `json:"refresh_token"`
// 			AccessToken  string `json:"access_token"`
// 		}

// 		var response Response
// 		json.Unmarshal(rr.Body.Bytes(), &response)

// 		accessToken = response.AccessToken
// 		refreshToken = response.RefreshToken

// 		var responseBody any
// 		json.NewDecoder(rr.Body).Decode(&responseBody)
// 		x, _ := json.Marshal(responseBody)
// 		fmt.Printf(">> login otp finish response code is: %d, output is: %v \n", rr.Code, string(x))
// 	}

// 	// check user
// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "admin@mail.com")
// 		printUser("after login otp", user, "admin1234", "1234", "")
// 	}

// 	// users get all
// 	{

// 		req := httptest.NewRequest("GET", "/users", nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		var responseBody any
// 		json.NewDecoder(rr.Body).Decode(&responseBody)
// 		x, _ := json.Marshal(responseBody)
// 		fmt.Printf(">> get all user response code is: %d, output is: %v \n", rr.Code, string(x))
// 	}

// 	// users get me
// 	{

// 		req := httptest.NewRequest("GET", "/users/me", nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		var responseBody any
// 		json.NewDecoder(rr.Body).Decode(&responseBody)
// 		x, _ := json.Marshal(responseBody)
// 		fmt.Printf(">> get me user response code is: %d, output is: %v \n", rr.Code, string(x))
// 	}

// 	// refresh token
// 	{

// 		body := struct {
// 			RefreshToken string `json:"refresh_token"`
// 		}{
// 			RefreshToken: refreshToken,
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/auth/refresh-token", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		var responseBody any
// 		json.NewDecoder(rr.Body).Decode(&responseBody)
// 		x, _ := json.Marshal(responseBody)
// 		fmt.Printf(">> refresh token response code is: %d, output is: %v \n", rr.Code, string(x))
// 	}

// 	// check user
// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "admin@mail.com")
// 		printUser("before password change request", user, "admin1234", "1234", "")
// 	}

// 	// password change request
// 	{

// 		req := httptest.NewRequest("POST", "/password/change/initiate", nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> password change initiate response code is: %d\n", rr.Code)
// 	}

// 	// check user
// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "admin@mail.com")
// 		printUser("after password change request", user, "admin1234", "1234", "123456")
// 	}

// 	// change password submit
// 	{

// 		body := struct {
// 			OTP         string `json:"otp"`
// 			OldPassword string `json:"old_password"`
// 			NewPassword string `json:"new_password"`
// 		}{
// 			OTP:         "123456",
// 			OldPassword: "admin1234",
// 			NewPassword: "4321admin",
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/password/change/verify", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> change password submit response code is: %d\n", rr.Code)
// 	}

// 	// check user
// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "admin@mail.com")
// 		printUser("after password change submit", user, "4321admin", "1234", "")
// 	}

// 	// pin change request
// 	{

// 		req := httptest.NewRequest("POST", "/pin/change/initiate", nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> pin change initiate response code is: %d\n", rr.Code)
// 	}

// 	// check user
// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "admin@mail.com")
// 		printUser("after pin change request", user, "4321admin", "1234", "123456")
// 	}

// 	// change pin submit
// 	{

// 		body := struct {
// 			OTP    string `json:"otp"`
// 			NewPIN string `json:"new_pin"`
// 		}{
// 			OTP:    "123456",
// 			NewPIN: "4321",
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/pin/change/verify", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> change pin submit response code is: %d\n", rr.Code)
// 	}

// 	// check user
// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "admin@mail.com")
// 		printUser("after pin change submit", user, "4321admin", "4321", "")
// 	}

// }
