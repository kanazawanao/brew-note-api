package api

import (
	"tripig/test/e2e/constant"

	client "github.com/tripig/api"
)

func TripigClient() *client.APIClient {
	conf := client.NewConfiguration()
	conf.Servers = client.ServerConfigurations{{
		URL: constant.ServerURL,
	}}

	client := client.NewAPIClient(conf)

	return client
}
