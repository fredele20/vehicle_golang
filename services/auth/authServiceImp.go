package auth

import (
	"context"
	"vehicle_golang/config"
	"vehicle_golang/models"
	repository "vehicle_golang/repository/user"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type AuthServiceImp struct {
	db     *mgo.Session
	repo   repository.UserRepo
	config *config.Configuration
}

func New(db *mgo.Session, c *config.Configuration) AuthService {
	return &AuthServiceImp{db: db, config: c, repo: repository.New(db, c)}
}

func (service *AuthServiceImp) Create(ctx context.Context, user *models.User) error {
	return service.repo.Create(ctx, user)
}

func (service *AuthServiceImp) Login(ctx context.Context, credential *models.LoginDetails) (*models.User, error) {
	query := bson.M{"email": credential.Email}
	user, err := service.repo.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(credential.Password); err != nil {
		return nil, err
	}

	return user, nil
}

// IsUserAlreadyExists checks if the user is already existing in db
func (service *AuthServiceImp) IsUserAlreadyExists(ctx context.Context, email string) bool {
	return service.repo.IsUserAlreadyExist(ctx, email)
}
