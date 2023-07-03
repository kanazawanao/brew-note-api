package test

import (
	"context"
	"net/http"
	"testing"
	"tripig/test/e2e/database"
	"tripig/test/e2e/tripig/api"

	"github.com/stretchr/testify/assert"
)

func TestTripigUsers(t *testing.T) {
	assert := assert.New(t)

	database.TruncateTable()
	database.ImportData()

	userApi := api.TripigClient().UserApi

	t.Run("Get User", func(t *testing.T) {
		result, res, _ := userApi.UsersGet(context.Background()).Execute()
		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Len(result, 5)
	})
}