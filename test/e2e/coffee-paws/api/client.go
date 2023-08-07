package api

import (
	"coffee-paws/test/e2e/constant"

	client "github.com/coffee-paws/api"
)

func CoffeePawsClient() *client.APIClient {
	conf := client.NewConfiguration()
	conf.Servers = client.ServerConfigurations{{
		URL: constant.ServerURL,
	}}

	client := client.NewAPIClient(conf)

	return client
}
