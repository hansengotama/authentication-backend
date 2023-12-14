package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/hansengotama/authentication-backend/internal/domain"
	sqlorder "github.com/hansengotama/authentication-backend/internal/lib/sql"
)

type GetOTPAuthParam struct {
	UserID int
	OTP    int
	Status domain.OTPAuthStatusEnum
	Order  sqlorder.SQLOrder
}

type OTPAuthGetRepositoryInterface interface {
	GetOTPAuth(context.Context, GetOTPAuthParam) (domain.OtpAuth, error)
}

type OTPAuthGetDB struct {
	postgresConn *sql.DB
}

func NewOTPAuthGetDB(postgresConn *sql.DB) OTPAuthGetRepositoryInterface {
	return OTPAuthGetDB{
		postgresConn: postgresConn,
	}
}

var ErrGetOTPAuthNotFound = errors.New("OTP Auth Not Found")

func (s OTPAuthGetDB) GetOTPAuth(ctx context.Context, param GetOTPAuthParam) (domain.OtpAuth, error) {
	query := "SELECT id, user_id, otp, otp_expired_at, status, created_at, updated_at from otp_auth"
	str := " WHERE "

	var params []any
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

	if param.Status.IsValid() {
		params = append(params, param.Status.String())
		query += str + fmt.Sprintf("status = $%d", len(params))
		str = " AND "
	}

	if param.Order.IsValid() {
		query += fmt.Sprintf(" ORDER BY %s %s", param.Order.Column, param.Order.By.String())
	}

	query += " LIMIT 1"

	rows, err := s.postgresConn.Query(query, params...)
	if err != nil {
		// logging
		return domain.OtpAuth{}, err
	}

	var otpAuth domain.OtpAuth
	var status string
	for rows.Next() {
		err := rows.Scan(&otpAuth.ID, &otpAuth.UserID, &otpAuth.OTP, &otpAuth.OTPExpiredAt, &status, &otpAuth.CreatedAt, &otpAuth.UpdatedAt)

		otpAuthStatus, err := domain.StringToOTPAuthStatus(status)
		otpAuth.Status = otpAuthStatus

		if err != nil {
			// logging
			return domain.OtpAuth{}, err
		}
	}

	isNotFound := otpAuth.ID == 0
	if isNotFound {
		return domain.OtpAuth{}, ErrGetOTPAuthNotFound
	}

	return otpAuth, nil
}
