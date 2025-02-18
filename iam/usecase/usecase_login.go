package usecase

// import (
// 	"context"
// 	"shared/core"
// 	ory "shared/helper/ory"
// )

// type LoginReq struct {
// 	Challenge string
// 	FlowID    string
// }

// type LoginRes struct{}

// type Login = core.ActionHandler[LoginReq, LoginRes]

// func ImplLogin(
// 	ory *ory.ORYServer,
// ) Login {
// 	return func(ctx context.Context, req LoginReq) (*LoginRes, error) {
// 		// redirect to login page if there is no login challenge or flow id in url query parameters
// 		if req.Challenge == "" && req.FlowID == "" {
// 			log.Println("No login challenge found or flow ID found in URL Query Parameters")

// 			// create oauth2 state and store in session
// 			b := make([]byte, 32)
// 			_, err := rand.Read(b)
// 			if err != nil {
// 				log.Errorf("generate state failed: %v", err)
// 				return &LoginRes{}, nil
// 			}
// 			state := base64.StdEncoding.EncodeToString(b)
// 			setSessionValue(w, r, "oauth2State", state)

// 			// start oauth2 authorization code flow
// 			redirectTo := ory.OAuth2Config.AuthCodeURL(state)
// 			log.Infof("redirect to hydra, url: %s", redirectTo)
// 			http.Redirect(w, r, redirectTo, http.StatusFound)
// 			return &LoginRes{}, nil
// 		}

// 		// if there is no flow id in url query parameters, create a new flow
// 		if req.FlowID == "" {
// 			// build url with hydra login challenge as url query parameter
// 			// it will be automatically passed to hydra upon redirect
// 			params := url.Values{
// 				"login_challenge": []string{req.Challenge},
// 			}
// 			redirectTo := fmt.Sprintf("%s/self-service/login/browser?", s.KratosPublicEndpoint) + params.Encode()
// 			http.Redirect(w, r, redirectTo, http.StatusFound)
// 			return &LoginRes{}, nil
// 		}

// 		// get cookie from headers
// 		cookie := r.Header.Get("cookie")
// 		// get the login flow
// 		flow, _, err := s.KratosAPIClient.FrontendApi.GetLoginFlow(r.Context()).Id(req.lowID).Cookie(cookie).Execute()
// 		if err != nil {
// 			writeError(w, http.StatusUnauthorized, err)
// 			return &LoginRes{}, nil
// 		}
// 		templateData := templateData{
// 			Title: "Login",
// 			UI:    &flow.Ui,
// 		}

// 		// render template index.html
// 		templateData.Render(w)
// 		return &LoginRes{}, nil
// 	}
// }
