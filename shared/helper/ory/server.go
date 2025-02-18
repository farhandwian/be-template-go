package ory

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"

	ory "github.com/ory/client-go"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

//go:embed config/idp.yml
var idpConfYAML []byte

type idpConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// ORYServer contains ORYServer information
type ORYServer struct {
	KratosAPIClient      *ory.APIClient
	KratosPublicEndpoint string
	HydraAPIClient       *ory.APIClient
	OAuth2Config         *oauth2.Config
	IDPConfig            *idpConfig
}

func NewServer(kratosPublicEndpointPort, hydraPublicEndpointPort, hydraAdminEndpointPort int) (*ORYServer, error) {
	// create a new kratos client for self hosted ORYServer
	conf := ory.NewConfiguration()
	conf.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:%d", kratosPublicEndpointPort)}}
	cj, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	conf.HTTPClient = &http.Client{Jar: cj}

	hydraConf := ory.NewConfiguration()
	hydraConf.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:%d", hydraAdminEndpointPort)}}

	idpConf := idpConfig{
		ClientID:     os.Getenv("CLIENT_ID"),     // ClientID:
		ClientSecret: os.Getenv("CLIENT_SECRET"), // Client
	}

	if err := yaml.Unmarshal(idpConfYAML, &idpConf); err != nil {
		return nil, err
	}

	oauth2Conf := &oauth2.Config{
		ClientID:     idpConf.ClientID,
		ClientSecret: idpConf.ClientSecret,
		RedirectURL:  fmt.Sprintf("http://localhost:%d/dashboard", os.Getenv("PORT")),
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("http://localhost:%d/oauth2/auth", hydraPublicEndpointPort),  // access from browser
			TokenURL: fmt.Sprintf("http://localhost:%d/oauth2/token", hydraPublicEndpointPort), // access from ORYServer
		},
		Scopes: []string{"openid", "offline"},
	}

	log.Println("OAuth2 Config: ", oauth2Conf)

	return &ORYServer{
		KratosAPIClient:      ory.NewAPIClient(conf),
		KratosPublicEndpoint: fmt.Sprintf("http://localhost:%d", kratosPublicEndpointPort),
		HydraAPIClient:       ory.NewAPIClient(hydraConf),
		OAuth2Config:         oauth2Conf,
		IDPConfig:            &idpConf,
	}, nil
}
