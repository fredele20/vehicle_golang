package user

import (
	"context"
	"vehicle_golang/models"
)

// UserRepo, used to perform DB operation
// Interface contain basic operation on user document
// So that, db operations can be performed easily
type UserRepo interface {
	// Create will perform db operation that will save the user
	// Returns error if there is any
	Create(context.Context, *models.User) error

	// FindAll returns all existing user in the database
	// Returns error if occured
	FindAll(context.Context) ([]*models.User, error)

	// FindById finds and returns the user with the provided id
	// Returns matched user and error if occured
	FindById(context.Context, string) (*models.User, error)

	// Update will update the user with the provided id
	// Returns an error if any
	Update(context.Context, interface{}, interface{}) error

	// Delete will remove user from the database
	// Returns error if any
	Delete(context.Context, *models.User) error

	// FindOne will find and return user with the query object
	// Query object is of type interface that can accept any object
	// returns matched user and error if any
	FindOne(context.Context, interface{}) (*models.User, error)

	// IsUserAlreadyExist check if a user is already existing in the db
	// returns true or false
	IsUserAlreadyExist(context.Context, string) bool
}
