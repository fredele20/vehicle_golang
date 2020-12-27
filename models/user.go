package models

import (
	"fmt"
	"vehicle_golang/utils"

	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name,omitempty" bson:"name,omitempty"`
	Surname  string        `json:"surname,omitempty" bson:"surname,omitempty"`
	Email    string        `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string        `json:"phone,omitempty" bson:"phone,omitempty"`
	Location *Location     `json:"location,omitempty" bson:"location,omitempty`
	Role     string        `json:"isAdmin,omitempty" bson:"isAdmin,omitempty"`
	Password string        `json:"password,omitempty" bson:"password,omitempty`
}

type Location struct {
	Longitude string `json:"longitude,omitempty" bson:"longitude,omitempty`
	Latitude  string `json:"latitude,omitempty" bson:"latitude,omitempty`
}

type LoginDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	fmt.Println("before storing: ", passwordHash)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)
	fmt.Println("after storing: ", u.Password)
	return nil
}

func (u *User) Initialize() error {
	salt := uuid.New().String()
	passwordByte := []byte(u.Password + salt)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword[:])
	u.Role = utils.UserRole

	return nil
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) Validate() error {
	if e := utils.ValidateRequiredAndLengthAndRegex(
		u.Name,
		true,
		3,
		255,
		"",
		"name",
	); e != nil {
		return e
	}

	if e := utils.ValidateRequiredAndLengthAndRegex(
		u.Surname,
		true,
		3,
		255,
		"",
		"surname",
	); e != nil {
		return e
	}

	if e := utils.ValidateRequiredAndLengthAndRegex(
		u.Email,
		true,
		3,
		255,
		"^[a-zA-z0-9.!#$%&'*+/?^_`{|}~]+@[a-zA-z0-9](?:[a-zA-z0-9]{0,"+
			"61}[a-zA-z0-9])?(?:\\.[a-zA-z0-9](?:[a-zA-z0-9]{0,61}[a-zA-z0-9])?)*$",
		"email",
	); e != nil {
		return e
	}

	if e := utils.ValidateRequiredAndLengthAndRegex(
		u.Phone,
		true,
		0,
		12,
		"",
		"phone",
	); e != nil {
		return e
	}

	if e := utils.ValidateRequiredAndLengthAndRegex(
		u.Location.Latitude,
		true,
		3,
		255,
		"",
		"location",
	); e != nil {
		return e
	}

	if e := utils.ValidateRequiredAndLengthAndRegex(
		u.Password,
		true,
		6,
		255,
		"",
		"password",
	); e != nil {
		return e
	}

	return nil
}
