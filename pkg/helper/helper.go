package helper

import (
	"crypto/rand"
	"math/big"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(p), err
}

func CompareHashedPassword(dbPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
}

func GenerateOTP() (string, error) {
	const digits = "0123456789"
	const otpLength = 6

	otp := make([]byte, otpLength)
	for i := range otp {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[n.Int64()]
	}

	return string(otp), nil
}

func CurrentIsoDateTimeString() string {
	return time.Now().Format(time.RFC3339)
}
