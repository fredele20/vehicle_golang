package user

import (
	"context"
	"vehicle_golang/config"
	"vehicle_golang/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type UserRepoImp struct {
	db     *mgo.Session
	config *config.Configuration
}

func New(db *mgo.Session, c *config.Configuration) UserRepo {
	return &UserRepoImp{db: db, config: c}
}

func (s *UserRepoImp) collection() *mgo.Collection {
	return s.db.DB(s.config.DatabaseName).C("users")
}

func (s *UserRepoImp) Create(ctx context.Context, user *models.User) error {
	return s.collection().Insert(user)
}

func (s *UserRepoImp) FindAll(ctx context.Context) ([]*models.User, error) {
	panic("Implement me")
}

func (s *UserRepoImp) Update(ctx context.Context, query, change interface{}) error {
	return s.collection().Update(query, change)
}

func (s *UserRepoImp) FindById(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := s.collection().Find(query).Select(bson.M{"password": 0, "salt": 0}).One(&user)
	return &user, err
}

func (s *UserRepoImp) FindOne(ctx context.Context, query interface{}) (*models.User, error) {
	var user models.User
	err := s.collection().Find(query).One(&user)
	return &user, err
}

func (s *UserRepoImp) IsUserAlreadyExist(ctx context.Context, email string) bool {
	query := bson.M{"email": email}
	_, err := s.FindOne(ctx, query)
	if err != nil {
		return false
	}

	return true
}

func (s *UserRepoImp) Delete(ctx context.Context, user *models.User) error {
	panic("implement me")
}
