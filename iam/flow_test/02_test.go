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
// 	"time"

// 	"github.com/joho/godotenv"
// )

// func Test02(t *testing.T) {
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

// 	// admin login
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

// 	accessToken := ""

// 	// admin login otp
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

// 		fmt.Printf(">> login otp finish with response code %d\n", rr.Code)

// 		type Response struct {
// 			RefreshToken string `json:"refresh_token"`
// 			AccessToken  string `json:"access_token"`
// 		}

// 		var response Response
// 		json.Unmarshal(rr.Body.Bytes(), &response)

// 		accessToken = response.AccessToken
// 	}

// 	// admin users get all
// 	{

// 		req := httptest.NewRequest("GET", "/users", nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		type ResponseBody struct {
// 			Count uint `json:"count"`
// 			Items any  `json:"items"`
// 		}

// 		var responseBody ResponseBody
// 		json.Unmarshal(rr.Body.Bytes(), &responseBody)

// 		fmt.Printf(">> get all user response code is: %d, total user is: %d \n", rr.Code, responseBody.Count)
// 	}

// 	// admin register one user
// 	{
// 		body := struct {
// 			Name        string            `json:"name"`
// 			Email       model.Email       `json:"email"`
// 			PhoneNumber model.PhoneNumber `json:"phone_number"`
// 		}{
// 			Name:        "user",
// 			Email:       model.Email("user@mail.com"),
// 			PhoneNumber: model.PhoneNumber("0897654321"),
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/account/register", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> login otp finish with response code %d\n", rr.Code)
// 	}

// 	userID := ""

// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "user@mail.com")
// 		printUser("before admin email activation request", user, "", "", "")

// 		userID = string(user.ID)
// 	}

// 	activationToken := ""

// 	// admin email activation request
// 	{
// 		body := struct {
// 			UserID model.UserID `json:"user_id"`
// 		}{
// 			UserID: model.UserID(userID),
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/account/activate/initiate", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> admin email activation request finish with response code %d %s\n", rr.Code, rr.Body.String())

// 		{
// 			userTokenPayloadInfo, _ := json.Marshal(model.UserTokenPayload{
// 				UserID:  body.UserID,
// 				Subject: model.EMAIL_ACTIVATION,
// 			})

// 			activationToken, _ = jwtToken.CreateToken(userTokenPayloadInfo, time.Now(), 1*time.Hour)
// 		}
// 	}

// 	// email activation submit
// 	{
// 		body := struct {
// 			ActivationToken string `json:"activation_token"`
// 			Password        string `json:"password"`
// 			Pin             string `json:"pin"`
// 		}{
// 			ActivationToken: activationToken,
// 			Password:        "112233",
// 			Pin:             "1111",
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/account/activate/verify", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> email activation submit finish with response code %d\n", rr.Code)
// 	}

// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "user@mail.com")
// 		printUser("after email activation submit", user, "112233", "1111", "")
// 	}

// 	// admin get users with id
// 	{

// 		req := httptest.NewRequest("GET", "/users/"+userID, nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		var responseBody any
// 		json.NewDecoder(rr.Body).Decode(&responseBody)
// 		x, _ := json.Marshal(responseBody)
// 		fmt.Printf(">> admin get users with id response code is: %d, output is: %v \n", rr.Code, string(x))

// 	}

// 	// users get access
// 	{

// 		req := httptest.NewRequest("GET", "/users/"+userID+"/access", nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		var responseBody any
// 		json.NewDecoder(rr.Body).Decode(&responseBody)
// 		x, _ := json.Marshal(responseBody)
// 		fmt.Printf(">> get user access response code is: %d, output is: %v \n", rr.Code, string(x))
// 	}

// 	resetPasswordToken := ""

// 	// password reset request
// 	{
// 		body := struct {
// 			UserID model.UserID `json:"user_id"`
// 		}{
// 			UserID: model.UserID(userID),
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/password/reset/initiate", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+accessToken)

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> password reset request finish with response code %d\n", rr.Code)

// 		{
// 			userTokenPayloadInfo, _ := json.Marshal(model.UserTokenPayload{
// 				UserID:  body.UserID,
// 				Subject: model.PASSWORD_RESET,
// 			})

// 			resetPasswordToken, _ = jwtToken.CreateToken(userTokenPayloadInfo, time.Now(), 1*time.Hour)
// 		}
// 	}

// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "user@mail.com")
// 		printUser("after email activation submit", user, "112233", "1111", "")
// 	}

// 	// password reset submit
// 	{
// 		body := struct {
// 			PasswordResetToken string `json:"password_reset_token"`
// 			NewPassword        string `json:"new_password"`
// 		}{
// 			PasswordResetToken: resetPasswordToken,
// 			NewPassword:        "22222",
// 		}

// 		bodyBytes, err := json.Marshal(body)
// 		if err != nil {
// 			t.Fatalf("Failed to marshal request body: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/password/reset/verify", bytes.NewBuffer(bodyBytes))
// 		req.Header.Set("Content-Type", "application/json")

// 		rr := httptest.NewRecorder()

// 		mux.ServeHTTP(rr, req)

// 		fmt.Printf(">> password reset submit finish with response code %d\n", rr.Code)
// 	}

// 	{
// 		var user model.User
// 		db.Find(&user, "email = ?", "user@mail.com")
// 		printUser("after email activation submit", user, "22222", "1111", "")
// 	}

// }
