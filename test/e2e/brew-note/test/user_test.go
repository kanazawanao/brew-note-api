package test

import (
	"brew-note/test/e2e/brew-note/api"
	"brew-note/test/e2e/database"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrewNoteUsers(t *testing.T) {
	assert := assert.New(t)

	database.TruncateTable()
	database.ImportUserData()

	brewNoteApi := api.BrewNoteClient().BrewNtoeApi

	t.Run("Get Users", func(t *testing.T) {
		result, response, _ := brewNoteApi.GetUsers(context.Background()).Execute()
		assert.Equal(http.StatusOK, response.StatusCode)
		assert.Equal(5, len(result))
	})
}
