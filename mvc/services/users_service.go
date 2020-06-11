package services

import (
	"application/mvc/domain"
	"application/mvc/util"
)

// GetUser is a model data base user
func GetUser(userID float64) (*domain.User, *util.ApplicationError) {
	return domain.GetUser(userID)
}
