package e

import "errors"

var (
	// User
	ErrUserNotFound             = errors.New("user not found")
	ErrEmailConflict            = errors.New("user with email already exist")
	ErrUsernameConflict         = errors.New("user with username already exist")
	ErrEmailAndUsernameConflict = errors.New("user with email and username already exist")
	ErrInvalidOtp               = errors.New("otp is invalid")
	ErrOtpTimeOut               = errors.New("your otp has timeout")

	ErrAdminNotFound = errors.New("user not found")

	ErrKeyNotFound = errors.New("given key not found in map")

	ErrNoRecordsAffected = errors.New("no records were affected")
)
