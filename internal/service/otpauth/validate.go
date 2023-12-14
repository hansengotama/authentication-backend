package otpauth

import (
	"context"
	"database/sql"
	"errors"
	db "github.com/hansengotama/authentication-backend/internal/repository/db/otpauth"
)

type ValidateOTPReq struct {
	UserID int `json:"user_id"`
	OTP    int `json:"otp"`
}

type ValidateOTPRes struct {
	UserID int `json:"user_id"`
}

type OtpAuthValidateServiceInterface interface {
	Validate(context.Context, ValidateOTPReq) (ValidateOTPRes, error)
}

type OtpAuthValidateService struct {
	postgresConn *sql.DB
}

func NewAuthValidateService(postgresConn *sql.DB) OtpAuthValidateServiceInterface {
	return OtpAuthValidateService{
		postgresConn: postgresConn,
	}
}

func (s OtpAuthValidateService) Validate(ctx context.Context, req ValidateOTPReq) (ValidateOTPRes, error) {
	// get by otp & expired at
	otpAuthGetDB := db.NewOTPAuthGetDB(s.postgresConn)
	res, err := otpAuthGetDB.Get(ctx, db.GetParam{
		UserID:  req.UserID,
		OTP:     req.OTP,
		OrderBy: "DESC",
	})
	if errors.Is(err, sql.ErrNoRows) {
		return ValidateOTPRes{}, err
	}

	if err != nil {
		return ValidateOTPRes{}, err
	}

	return ValidateOTPRes{
		UserID: res.ID,
	}, nil
}
