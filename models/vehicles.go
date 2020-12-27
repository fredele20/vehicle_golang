package models

import (
	"vehicle_golang/utils"

	"github.com/globalsign/mgo/bson"
)

type Vehicle struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"id,omitempty"`
	Name        string        `json:"name,omitempty" bson:"name,omitempty"`
	Brand       string        `json:"brand,omitempty" bson:"brand,omitempty"`
	PlateNumber string        `json:"plateNumber,omitempty" bson:"plateNumber,omitempty"`
	IsOwned     bool          `json:"isOwned,omitempty" bson:"isOwned,omitempty"`
}

func (v *Vehicle) Validate() error {
	if e := utils.ValidateRequiredAndLengthAndRegex(
		v.Name,
		true,
		0,
		0,
		"",
		"name",
	); e != nil {
		return e
	}

	if e := utils.ValidateRequiredAndLengthAndRegex(
		v.Brand,
		true,
		0,
		0,
		"",
		"brand",
	); e != nil {
		return e
	}

	if e := utils.ValidateRequiredAndLengthAndRegex(
		v.PlateNumber,
		true,
		0,
		0,
		"",
		"plateNumber",
	); e != nil {
		return e
	}

	return nil
}
