package constants

type AccType string

const (
	AccountPublic  AccType = "PUBLIC"
	AccountPrivate AccType = "PRIVATE"
)

type UserStatus string

const (
	UserStatusActive  UserStatus = "ACTIVE"
	UserStatusPending UserStatus = "PENDING"
	UserStatusBlocked UserStatus = "BLOCKED"
)
