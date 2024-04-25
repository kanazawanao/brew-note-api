package api

import (
	"brew-note/test/e2e/constant"

	client "github.com/brew-note/api"
)

func BrewNoteClient() *client.APIClient {
	conf := client.NewConfiguration()
	conf.Servers = client.ServerConfigurations{{
		URL: constant.ServerURL,
	}}

	client := client.NewAPIClient(conf)

	return client
}
