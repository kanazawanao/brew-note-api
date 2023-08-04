package test

import (
	"context"
	"net/http"
	"testing"
	"tripig/test/e2e/data"
	"tripig/test/e2e/database"
	"tripig/test/e2e/tripig/api"

	"github.com/stretchr/testify/assert"
)

func TestTripigUsers(t *testing.T) {
	assert := assert.New(t)

	database.TruncateTable()
	database.ImportUserData()

	tripigApi := api.TripigClient().TripigApi

	t.Run("Get Users", func(t *testing.T) {
		result, response, _ := tripigApi.GetUsers(context.Background()).Execute()
		assert.Equal(http.StatusOK, response.StatusCode)
		assert.Equal(5, len(result))
	})
}

func TestTripigUser(t *testing.T) {
	assert := assert.New(t)

	database.TruncateTable()
	database.ImportUserData()

	tripigApi := api.TripigClient().TripigApi

	t.Run("Get User", func(t *testing.T) {
		result, response, _ := tripigApi.GetUser(context.Background(), data.UserIDs[0]).Execute()
		assert.Equal(http.StatusOK, response.StatusCode)
		assert.Equal(data.UserIDs[0], *result.Id)
	})
}