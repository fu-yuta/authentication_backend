package errorValues

import (
	"errors"
)

var (
	AlreadyExistUserError  = errors.New("already exit this user")
	CreateUserError        = errors.New("Failed Create New User")
	NotRegisterUserError   = errors.New("Not Registered User")
	MissmatchPasswordError = errors.New("Password Missmatch")
	LoginError             = errors.New("Failed Login")
	AccessTokenExpireError = errors.New("AccessToken is Expire")
)
