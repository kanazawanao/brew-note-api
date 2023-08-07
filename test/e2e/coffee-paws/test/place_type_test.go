package test

import (
	"coffee-paws/test/e2e/coffee-paws/api"
	"coffee-paws/test/e2e/database"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoffeePawsPlaceType(t *testing.T) {
	assert := assert.New(t)

	database.TruncateTable()
	database.ImportUserData()

	coffeePawsApi := api.CoffeePawsClient().CoffeePawsApi

	t.Run("Get Place Types", func(t *testing.T) {
		result, response, _ := coffeePawsApi.GetPlaceTypes(context.Background()).Execute()
		assert.Equal(http.StatusOK, response.StatusCode)
		assert.Equal(5, len(result))
	})
}
