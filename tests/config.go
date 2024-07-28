package tests

import (
	"github.com/go-resty/resty/v2"
)

type API struct {
	Client *resty.Client
}

func SetupHeadersAgente() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func SetupApi() *API {

	api := resty.New().
		SetBaseURL("http://localhost:8080")

	return &API{
		Client: api,
	}

}
