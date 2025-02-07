package controller

import (
	"encoding/json"
	"fmt"
	"iam/model"
	"net/http"
	"shared/core"
	"shared/helper"
	"time"
)

func ExperimentHandler(mux *http.ServeMux) helper.APIData {

	type Body struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Status   string `json:"status"`
	}

	type UseCaseReq struct {
		Body       Body          `json:"body" http:"body"`
		ID         string        `json:"id" http:"path"`
		Page       int           `json:"page" http:"query"`
		Size       int           `json:"size" http:"query"`
		Keyword    string        `json:"keyword" http:"query"`
		UserAccess string        `json:"access" http:"context"`
		Now        time.Time     `json:"now" http:"now"`
		OTP        time.Duration `json:"otp" http:"func(otp)"`
	}

	apiData := helper.APIData{
		Access:  model.DEFAULT_OPERATION,
		Method:  http.MethodPost,
		Url:     "/hello/{id}",
		Body:    Body{},
		Summary: "Experiment",
		Tag:     "Experiment",
		QueryParams: []helper.QueryParam{
			{Name: "page", Type: "integer", Description: "Page number", Required: false},
			{Name: "size", Type: "integer", Description: "Number of items per page", Required: false},
			{Name: "keyword", Type: "string", Description: "any keyword", Required: false},
		},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		x, ok := ExtractRequest[UseCaseReq](w, r, apiData.Url, func(key string) (any, error) {

			if key == "otp" {
				return time.Duration(5 * time.Second), nil
			}

			return nil, nil
		})
		if !ok {
			return
		}

		a, _ := json.MarshalIndent(x, "", " ")
		fmt.Printf("result %v\n", string(a))

		format := "2006-01-02 15:04:05"

		var loadLocation string
		var errMsg string
		location, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			errMsg = err.Error()
			loadLocation = ""
		} else {
			loadLocation = time.Now().In(location).Format(format)
		}

		name, offset := time.Now().Zone()

		offsetHours := offset / 3600

		Success(w, map[string]any{
			"name":                   name,
			"offset_hours":           offsetHours,
			"now":                    time.Now().Format(format),
			"local":                  time.Now().Local().Format(format),
			"utc":                    time.Now().UTC().Format(format),
			"plus7":                  time.Now().Add(7 * time.Hour).Format(format),
			"load_location":          loadLocation,
			"has_error_LoadLocation": errMsg,
		})

	}

	const accessType core.ContextKey = "access"

	mid := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			ctx := core.AttachDataToContext[string](r.Context(), accessType, "admin")

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

		}
	}

	mux.HandleFunc(apiData.GetMethodUrl(), mid(handler))

	return apiData
}
