package test

import (
	"coffee-paws/test/e2e/coffee-paws/api"
	"coffee-paws/test/e2e/database"
	"context"
	"net/http"
	"testing"

	openapi "github.com/coffee-paws/api"
	"github.com/stretchr/testify/assert"
)

func TestCoffeePawsPostStore(t *testing.T) {
	assert := assert.New(t)

	database.TruncateTable()

	coffeePawsApi := api.CoffeePawsClient().CoffeePawsApi

	t.Run("Post Store", func(t *testing.T) {
		request := &openapi.CreateStore{}
		request.Name = "name"
		request.Address = "address"
		request.Url = "url"
		request.PlaceId = ""
		result, response, _ := coffeePawsApi.CreateStore(context.Background()).CreateStore(*request).Execute()

		assert.Equal(http.StatusOK, response.StatusCode)
		assert.NotEmpty(result.Name)
		assert.NotEmpty(result.Address)
	})

	t.Run("Get Stores", func(t *testing.T) {
		result, response, _ := coffeePawsApi.GetStores(context.Background()).Execute()
		assert.Equal(http.StatusOK, response.StatusCode)
		assert.Equal(1, len(result))
	})
}