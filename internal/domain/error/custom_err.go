package e

import "errors"

var (
	// User
	ErrUserNotFound             = errors.New("user not found")
	ErrEmailConflict            = errors.New("user with email already exist")
	ErrUsernameConflict         = errors.New("user with username already exist")
	ErrEmailAndUsernameConflict = errors.New("user with email and username already exist")
	
	ErrAdminNotFound             = errors.New("user not found")
	// ErrIsEmpty                  = errors.New("is empty")
	// ErrDb                       = errors.New("db error")
	// ErrInvalidPassword          = errors.New("invalid password")
	// ErrInvalidStatusValue       = errors.New("invalid status value")
	// ErrNotAvailable             = errors.New("not available")
	// ErrQuantityExceeds          = errors.New("selected quantity not available")
	// ErrInvalidCoupon            = errors.New("invalid coupon")
	// ErrCouponNotApplicable      = errors.New("coupon doesn't meet terms")
	// ErrCouponAlreadyRedeemed    = errors.New("coupon already redeemed")
)
