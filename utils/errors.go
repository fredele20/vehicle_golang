package utils

import (
	"errors"
	"regexp"
	"strconv"
)

// ValidateRequiredAndLengthAndRegex is used to validate any input data but in string type
// @params value is the input value
// @params isRequired is a bool that defines if the input value is required or not
// @params minLength is the minimum length of the input value, 0 value defines no min length checked
// @params maxLength is the maximum length of the input value, 0 value defines no max length checked
// regex defined the regex of the input value, "" value defines no regex required
func ValidateRequiredAndLengthAndRegex(value string, isRequired bool, minLength, maxLength int, regex, fieldName string) error {
	length := len(value)
	re := regexp.MustCompile(regex)

	if isRequired == true && length < 1 {
		return errors.New(fieldName + " is required")
	}

	// Min Length check
	if minLength != 0 && length > 1 && length < minLength {
		return errors.New(fieldName + " must be min " + strconv.Itoa(minLength))
	}

	// Max Length check
	if maxLength != 0 && length > 1 && length > maxLength {
		return errors.New(fieldName + " must not be more than " + strconv.Itoa(maxLength))
	}

	// Regex check
	if !re.MatchString(value) {
		return errors.New("Invalid " + fieldName)
	}
	return nil
}

var errorMessage = map[string]string{
	"invalidUserID":  "invalid user id",
	"internalError":  "an internal error occured",
	"userNotFound":   "user could not be found",
	"unauthenticaed": "access denied, unauthenticated",
	"unauthorized":   "access denied, Forbidden",
}

// NewHTTPError creates error model that will send as http response
// if any error occurs
func NewHTTPError(errorCode string, statusCode int) map[string]interface{} {
	m := make(map[string]interface{})
	m["error"] = errorCode
	m["error_description"], _ = errorMessage[errorCode]
	m["code"] = statusCode

	return m
}

// NewHTTPCustomError creates error model that will send as http response
// if any error occurs
func NewHTTPCustomError(errorCode, errorMsg string, statusCode int) map[string]interface{} {

	m := make(map[string]interface{})
	m["error"] = errorCode
	m["error_description"] = errorMsg
	m["code"] = statusCode

	return m
}

// Error codes
const (
	InvalidUserID       = "invalidUserID" // in case userid not exists
	InternalError       = "internalError" // in case, any internal server error occurs
	UserNotFound        = "userNotFound"  // if user not found
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "entityCreationError"
	Unauthenticated     = "unauthenticated"
	Unauthorized        = "unauthorized" // in case, try to access restricted resource
	BadRequest          = "badRequest"
	UserAlreadyExists   = "userAlreadyExists"
)
