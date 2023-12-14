package db

import (
	"context"
	"database/sql"
	"time"
)

type InsertOTPAuthParam struct {
	UserID       int
	OTP          int
	OTPExpiredAt time.Time
}

type InsertOTPAuthResponse struct {
	UserId       int
	OTP          int
	OTPExpiredAt time.Time
}

type OTPAuthInsertRepositoryInterface interface {
	Insert(context.Context, InsertOTPAuthParam) (InsertOTPAuthResponse, error)
}

type OTPAuthInsertDB struct {
	postgresConn *sql.DB
}

func NewOTPAuthInsertDB(postgresConn *sql.DB) OTPAuthInsertDB {
	return OTPAuthInsertDB{
		postgresConn: postgresConn,
	}
}

func (s OTPAuthInsertDB) Insert(ctx context.Context, param InsertOTPAuthParam) error {
	_, err := s.postgresConn.Exec("INSERT INTO otp_auth(user_id, otp, otp_expired_at) VALUES ($1, $2, $3)", param.UserID, param.OTP, param.OTPExpiredAt)
	if err != nil {
		// logging
		return err
	}

	return nil
}
