package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hansengotama/authentication-backend/internal/domain"
)

type GetParam struct {
	UserID  int
	OTP     int
	OrderBy string
}

type OTPAuthGetRepositoryInterface interface {
	Get(context.Context, GetParam) (domain.OtpAuth, error)
}

type OTPAuthGetDB struct {
	postgresConn *sql.DB
}

func NewOTPAuthGetDB(postgresConn *sql.DB) OTPAuthGetRepositoryInterface {
	return OTPAuthGetDB{
		postgresConn: postgresConn,
	}
}

func (s OTPAuthGetDB) Get(ctx context.Context, param GetParam) (domain.OtpAuth, error) {
	query := "SELECT id, user_id, otp, otp_expired_at, created_at, updated_at from otp_auth"
	str := " WHERE "

	params := []any{}
	if param.UserID != 0 {
		params = append(params, param.UserID)
		query += str + fmt.Sprintf("user_id = $%d", len(params))
		str = " AND "
	}

	if param.OTP != 0 {
		params = append(params, param.OTP)
		query += str + fmt.Sprintf("otp = $%d", len(params))
		str = " AND "
	}

	if param.OrderBy != "" {
		query += fmt.Sprintf(" ORDER BY created_at %s", param.OrderBy)
	}

	query += " LIMIT 1"

	rows, err := s.postgresConn.Query(query, params...)
	if err != nil {
		// logging
		return domain.OtpAuth{}, err
	}

	var otpAuth domain.OtpAuth
	for rows.Next() {
		err := rows.Scan(&otpAuth.ID, &otpAuth.UserID, &otpAuth.OTP, &otpAuth.OTPExpiredAt, &otpAuth.CreatedAt, &otpAuth.UpdatedAt)
		if err != nil {
			return domain.OtpAuth{}, err
		}
	}

	return otpAuth, nil
}
