package test

import (
	"context"
	"net/http"
	"testing"
	"tripig/test/e2e/database"
	"tripig/test/e2e/tripig/api"

	"github.com/stretchr/testify/assert"
)

func TestTripigPlaceType(t *testing.T) {
	assert := assert.New(t)

	database.TruncateTable()
	database.ImportUserData()

	tripigApi := api.TripigClient().TripigApi

	t.Run("Get Place Types", func(t *testing.T) {
		result, response, _ := tripigApi.GetPlaceTypes(context.Background()).Execute()
		assert.Equal(http.StatusOK, response.StatusCode)
		assert.Equal(5, len(result))
	})
}
