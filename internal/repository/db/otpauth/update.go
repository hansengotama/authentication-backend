package db

import (
	"context"
	"database/sql"
	"github.com/hansengotama/authentication-backend/internal/domain"
	"time"
)

type UpdateOTPAuthStatusParam struct {
	ID     int
	Status domain.OTPAuthStatusEnum
}

type OTPAuthUpdateRepositoryInterface interface {
	UpdateOTPAuthStatus(context.Context, UpdateOTPAuthStatusParam) error
}

type OTPAuthUpdateDB struct {
	postgresConn *sql.DB
}

func NewOTPAuthUpdateDB(postgresConn *sql.DB) OTPAuthUpdateDB {
	return OTPAuthUpdateDB{
		postgresConn: postgresConn,
	}
}

func (s OTPAuthUpdateDB) UpdateOTPAuthStatus(ctx context.Context, param UpdateOTPAuthStatusParam) error {
	_, err := s.postgresConn.Exec("UPDATE otp_auth SET status = $1, updated_at = $2 WHERE id = $3", param.Status.String(), time.Now(), param.ID)
	if err != nil {
		// logging
		return err
	}

	return nil
}
