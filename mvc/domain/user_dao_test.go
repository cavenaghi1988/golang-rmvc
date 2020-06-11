package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "We ware not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "User 0 not found", err.Meassage)
	assert.EqualValues(t, "no found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "John", user.FirstName)
	assert.EqualValues(t, "Smith", user.LastName)
	assert.EqualValues(t, "john@gmail.com", user.Email)
}
