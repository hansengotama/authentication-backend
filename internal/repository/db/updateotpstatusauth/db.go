package updateotpstatusauth

import (
	"context"
	"github.com/hansengotama/authentication-backend/internal/domain"
	"github.com/hansengotama/authentication-backend/internal/lib/postgres"
	"time"
)

type UpdateOTPAuthStatusParam struct {
	ID     int
	Status domain.OTPAuthStatusEnum
}

type UpdateOTPAuthStatusDBInterface interface {
	UpdateOTPAuthStatus(context.Context, postgres.SQLExecutor, UpdateOTPAuthStatusParam) error
}

type UpdateOTPAuthStatusDB struct{}

func (s UpdateOTPAuthStatusDB) UpdateOTPAuthStatus(ctx context.Context, executor postgres.SQLExecutor, param UpdateOTPAuthStatusParam) error {
	_, err := executor.ExecContext(ctx, "UPDATE otp_auth SET status = $1, updated_at = $2 WHERE id = $3", param.Status.String(), time.Now(), param.ID)
	if err != nil {
		// logging
		return err
	}

	return nil
}
