package domain

import "time"

type OtpAuth struct {
	ID           int
	UserID       int
	OTP          int
	OTPExpiredAt int

	CreatedAt time.Time
	UpdatedAt time.Time
}
