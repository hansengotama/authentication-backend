package db

import (
	"context"
	"database/sql"
	"github.com/hansengotama/authentication-backend/internal/domain"
	"time"
)

type InsertOTPAuthParam struct {
	UserID       int
	OTP          int
	Status       domain.OTPAuthStatusEnum
	OTPExpiredAt time.Time
}

type InsertOTPAuthResponse struct {
	UserId       int
	OTP          int
	OTPExpiredAt time.Time
}

type OTPAuthInsertRepositoryInterface interface {
	InsertOTPAuth(context.Context, InsertOTPAuthParam) error
}

type OTPAuthInsertDB struct {
	postgresConn *sql.DB
}

func NewOTPAuthInsertDB(postgresConn *sql.DB) OTPAuthInsertDB {
	return OTPAuthInsertDB{
		postgresConn: postgresConn,
	}
}

func (s OTPAuthInsertDB) InsertOTPAuth(ctx context.Context, param InsertOTPAuthParam) error {
	_, err := s.postgresConn.Exec("INSERT INTO otp_auth(user_id, otp, otp_expired_at, status) VALUES ($1, $2, $3, $4)", param.UserID, param.OTP, param.OTPExpiredAt, param.Status.String())
	if err != nil {
		// logging
		return err
	}

	return nil
}
