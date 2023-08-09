package test

import (
	"coffee-paws/test/e2e/coffee-paws/api"
	"coffee-paws/test/e2e/data"
	"coffee-paws/test/e2e/database"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoffeePawsUsers(t *testing.T) {
	assert := assert.New(t)

	database.TruncateTable()
	database.ImportUserData()

	coffeePawsApi := api.CoffeePawsClient().CoffeePawsApi

	t.Run("Get Users", func(t *testing.T) {
		result, response, _ := coffeePawsApi.GetUsers(context.Background()).Execute()
		assert.Equal(http.StatusOK, response.StatusCode)
		assert.Equal(5, len(result))
	})
	
	t.Run("Get User", func(t *testing.T) {
		result, response, _ := coffeePawsApi.GetUser(context.Background(), data.UserIDs[0]).Execute()
		assert.Equal(http.StatusOK, response.StatusCode)
		assert.Equal(data.UserIDs[0], *result.Id)
	})
}
