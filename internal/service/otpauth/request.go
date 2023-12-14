package otpauth

import (
	"context"
	"database/sql"
	"github.com/hansengotama/authentication-backend/internal/domain"
	"github.com/hansengotama/authentication-backend/internal/lib/env"
	"github.com/hansengotama/authentication-backend/internal/lib/generator"
	db "github.com/hansengotama/authentication-backend/internal/repository/db/otpauth"
	"time"
)

type RequestOTPReq struct {
	UserID int `json:"user_id"`
}

type RequestOTPQueryRes struct {
	UserID int
	OTP    int
}

type OtpAuthRequestServiceInterface interface {
	Request(context.Context, RequestOTPReq) (RequestOTPQueryRes, error)
}

type OtpAuthRequestService struct {
	postgresConn *sql.DB
}

func NewAuthRequestService(postgresConn *sql.DB) OtpAuthRequestServiceInterface {
	return OtpAuthRequestService{
		postgresConn: postgresConn,
	}
}

func (s OtpAuthRequestService) Request(ctx context.Context, req RequestOTPReq) (RequestOTPQueryRes, error) {
	otp, err := generator.RandomNumbers(5)
	if err != nil {
		return RequestOTPQueryRes{}, err
	}

	otpAuthInsertDB := db.NewOTPAuthInsertDB(s.postgresConn)
	param := db.InsertOTPAuthParam{
		UserID:       req.UserID,
		OTP:          otp,
		OTPExpiredAt: time.Now().Add(env.GetOTPExpirationTime()),
		Status:       domain.OTPAuthStatusEnumCreated,
	}
	err = otpAuthInsertDB.InsertOTPAuth(ctx, param)
	if err != nil {
		return RequestOTPQueryRes{}, err
	}

	return RequestOTPQueryRes{
		UserID: req.UserID,
		OTP:    otp,
	}, nil
}
