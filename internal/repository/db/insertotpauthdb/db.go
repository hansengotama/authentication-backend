package insertotpauthdb

import (
	"context"
	"github.com/hansengotama/authentication-backend/internal/domain"
	"github.com/hansengotama/authentication-backend/internal/lib/postgres"
	"time"
)

type InsertOTPAuthParam struct {
	UserID       int
	OTP          int
	Status       domain.OTPAuthStatusEnum
	OTPExpiredAt time.Time
}

type InsertOTPAuthDBInterface interface {
	InsertOTPAuth(context.Context, postgres.SQLExecutor, InsertOTPAuthParam) error
}

type InsertOTPAuthDB struct{}

func (s InsertOTPAuthDB) InsertOTPAuth(ctx context.Context, executor postgres.SQLExecutor, param InsertOTPAuthParam) error {
	_, err := executor.ExecContext(ctx, "INSERT INTO otp_auth(user_id, otp, otp_expired_at, status) VALUES ($1, $2, $3, $4)", param.UserID, param.OTP, param.OTPExpiredAt, param.Status.String())
	if err != nil {
		// logging
		return err
	}

	return nil
}
