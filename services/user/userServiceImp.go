package user

import (
	"context"
	"vehicle_golang/config"
	"vehicle_golang/models"
	repository "vehicle_golang/repository/user"

	"github.com/globalsign/mgo"
)

type UserServiceImp struct {
	db     *mgo.Session
	repo   repository.UserRepo
	config *config.Configuration
}

// New function will initialize the UserService
// taking db and session in params
// db session is required to perform db operations
// config is required to get the config info
func New(db *mgo.Session, c *config.Configuration) UserService {
	return &UserServiceImp{db: db, config: c, repo: repository.New(db, c)}
}

// Get function will find user by id
// returns user and error if any
func (s *UserServiceImp) Get(ctx context.Context, id string) (*models.User, error) {
	return s.repo.FindById(ctx, id)
}
