package db

import (
	"vehicle_golang/config"

	"github.com/globalsign/mgo"
)

var instance *mgo.Session

var err error

// GetInstance returns a copy of the db session
func GetInstance(c *config.Configuration) *mgo.Session {

	if instance == nil {
		instance, err = mgo.Dial(c.DatabaseConnectionURL)
		if err != nil {
			panic(err)
		}
	}

	return instance.Copy()
}
