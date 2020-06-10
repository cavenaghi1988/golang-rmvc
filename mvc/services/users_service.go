package services

import (
	"mvc/domain"
	"mvc/util"
)

// GetUser is a model data base user
func GetUser(userID float64) (*domain.User, *util.ApplicationError) {
	return domain.GetUser(userID)
}
