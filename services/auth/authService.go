package auth

import (
	"context"
	"vehicle_golang/models"
)

type AuthService interface {
	Create(context.Context, *models.User) error
	Login(context.Context, *models.LoginDetails) (*models.User, error)
	IsUserAlreadyExists(context.Context, string) bool
}
