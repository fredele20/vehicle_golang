package utils

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// UnixMilli used to get the millisecond of a given time
// @params t
// returns the millisecond of the given time
func UnixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / (int64(time.Millisecond)))
}

// CurrentTimeInMilliseconds use to get current time in milliseconds
// This function will be used when we need the current timestamp
// This function returns current timestamp in milliseconds
func CurrentTimeInMilliseconds() int64 {
	return UnixMilli(time.Now())
}

/***************************************** Bson Operation ******************************************/
//CustomBson Operation
// CustomBson used to perform custom bson related operations
// like set, push, unset e.t.c by using struct
// This will be very useful when we need to creat bson map for struct
type CustomBson struct{}

// BsonWrapper contains basic bson operations like $set, $push, $addToSet
// It is very useful to convert struct in bson
type BsonWrapper struct {
	// Set will set the data into the database
	// example - if it needs to set "name":"John", then it need to create a struct
	// that contains that name field and assign that struct in that field. After
	// encoded in bson it will look like { $set: { name: "John" } } and this
	// will be very useful in mongodb query
	Set interface{} `json:"$set,omitempty" bson:"$set,omitempty"`

	// The Unset operator delete a particular field
	// If the field does not exist, then Unset does nothing
	// If it needs to unset a name field then we simply create a struct that
	// contains name field and then set "" to name.
	// Now to unset, set that struct to Unset field. After encoded that it will
	// look like { $unset: { name: "" } }
	Unset interface{} `json:"$unset,omitempty" bson:"$unset,omitempty"`

	// The push appends a specified field to an array
	// If the field is absent in the document, it will update
	// Push adds the array field with the value as it's element.
	// If the field is not an array the operation will fail.
	Push interface{} `json:"$push,omitempty" bson:"$push,omitempty"`

	// The AddToSet operator adds a value to an array unless the value is already
	// present, in which case AddToSet does nothing to the array
	AddToSet interface{} `json:"$addToSet,omitempty" bson:"$addToSet,omitempty"`
}

// ToMap function converts interface to map
// It takes interface as param and returns map, or error if any.
func ToMap(i interface{}) (map[string]interface{}, error) {
	var stringInterfaceMap map[string]interface{}
	itr, _ := bson.Marshal(i)
	err := bson.Unmarshal(itr, &stringInterfaceMap)
	return stringInterfaceMap, err
}

// Set function creates a query to replace the value of a field with the specified value
// param - data that needs to be set
// returns - query map and error if any
func (c *CustomBson) Set(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Set: data}
	return ToMap(s)
}

// Push function creates a query to append a specified value to an array field
// param - data that needs to be set
// returns - query map and error if any
func (c *CustomBson) Push(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Push: data}
	return ToMap(s)
}

// Unset function creates a query to delete a particular field
// param - data that needs to be unset
// returns - query map and error if any
func (c *CustomBson) Unset(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Unset: data}
	return ToMap(s)
}

// AddToSet creates query to add a value to an array unless the value is already present.
// params - data that need to add to set,
// returns - query map and error if nay
func (customBson *CustomBson) AddToSet(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{AddToSet: data}
	return ToMap(s)
}

// End of Bson
