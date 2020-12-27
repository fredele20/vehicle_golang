package user

import (
	"context"
	"vehicle_golang/models"
)

type UserService interface {
	Get(context.Context, string) (*models.User, error)
}
