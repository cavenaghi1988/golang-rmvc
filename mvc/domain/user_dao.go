package domain

import (
	"fmt"
	"mvc/util"
	"net/http"
)

var (
	users = map[float64]*User{
		123: {ID: 1, FirstName: "John", LastName: "Smith", Email: "john@gmail.com"},
	}
)

func GetUser(userID float64) (*User, *util.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &util.ApplicationError{
		Meassage: fmt.Sprintf("User %v not found", userID),
		Status:   http.StatusNotFound,
		Code:     "no found",
	}
}
