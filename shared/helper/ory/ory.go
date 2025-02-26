package ory

import (
	"fmt"

	ory "github.com/ory/client-go"
)

func SetupOryClient() *ory.APIClient {
	c := ory.NewConfiguration()
	c.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:5000")}}
	oryApp := ory.NewAPIClient(c)
	return oryApp
}
