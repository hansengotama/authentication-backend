package domain

import (
	"fmt"
	"strings"
	"time"
)

type OtpAuth struct {
	ID           int
	UserID       int
	OTP          int
	Status       OTPAuthStatusEnum
	OTPExpiredAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

const OTPAuthStatusCreated = "created"
const OTPAuthStatusExpired = "expired"
const OTPAuthStatusValidated = "validated"

type OTPAuthStatusEnum int

const (
	OTPAuthStatusEnumCreated OTPAuthStatusEnum = iota
	OTPAuthStatusEnumExpired
	OTPAuthStatusEnumValidated
)

func (e OTPAuthStatusEnum) String() string {
	return [...]string{OTPAuthStatusCreated, OTPAuthStatusExpired, OTPAuthStatusValidated}[e]
}

func (e OTPAuthStatusEnum) IsValid() bool {
	return e == OTPAuthStatusEnumCreated ||
		e == OTPAuthStatusEnumExpired ||
		e == OTPAuthStatusEnumValidated
}

func StringToOTPAuthStatus(s string) (OTPAuthStatusEnum, error) {
	switch strings.ToLower(s) {
	case OTPAuthStatusCreated:
		return OTPAuthStatusEnumCreated, nil
	case OTPAuthStatusExpired:
		return OTPAuthStatusEnumExpired, nil
	case OTPAuthStatusValidated:
		return OTPAuthStatusEnumValidated, nil
	default:
		return -1, fmt.Errorf("invalid status: %s", s)
	}
}
