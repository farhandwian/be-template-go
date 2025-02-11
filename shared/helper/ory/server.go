package ory

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"

	ory "github.com/ory/client-go"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

//go:embed config/idp.yml
var idpConfYAML []byte

type idpConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Port         int    `yaml:"port"`
}

// server contains server information
type server struct {
	KratosAPIClient      *ory.APIClient
	KratosPublicEndpoint string
	HydraAPIClient       *ory.APIClient
	Port                 string
	OAuth2Config         *oauth2.Config
	IDPConfig            *idpConfig
}

func NewServer(kratosPublicEndpointPort, hydraPublicEndpointPort, hydraAdminEndpointPort int) (*server, error) {
	// create a new kratos client for self hosted server
	conf := ory.NewConfiguration()
	conf.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://kratos:%d", kratosPublicEndpointPort)}}
	cj, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	conf.HTTPClient = &http.Client{Jar: cj}

	hydraConf := ory.NewConfiguration()
	hydraConf.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://hydra:%d", hydraAdminEndpointPort)}}

	idpConf := idpConfig{}

	if err := yaml.Unmarshal(idpConfYAML, &idpConf); err != nil {
		return nil, err
	}

	oauth2Conf := &oauth2.Config{
		ClientID:     idpConf.ClientID,
		ClientSecret: idpConf.ClientSecret,
		RedirectURL:  fmt.Sprintf("http://localhost:%d/dashboard", idpConf.Port),
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("http://localhost:%d/oauth2/auth", hydraPublicEndpointPort), // access from browser
			TokenURL: fmt.Sprintf("http://hydra:%d/oauth2/token", hydraPublicEndpointPort),    // access from server
		},
		Scopes: []string{"openid", "offline"},
	}

	log.Println("OAuth2 Config: ", oauth2Conf)

	return &server{
		KratosAPIClient:      ory.NewAPIClient(conf),
		KratosPublicEndpoint: fmt.Sprintf("http://localhost:%d", kratosPublicEndpointPort),
		HydraAPIClient:       ory.NewAPIClient(hydraConf),
		Port:                 fmt.Sprintf(":%d", idpConf.Port),
		OAuth2Config:         oauth2Conf,
		IDPConfig:            &idpConf,
	}, nil
}
