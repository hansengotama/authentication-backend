package otpauth

import (
	"context"
	"database/sql"
	"github.com/hansengotama/authentication-backend/internal/domain"
	sqlorder "github.com/hansengotama/authentication-backend/internal/lib/sql"
	db "github.com/hansengotama/authentication-backend/internal/repository/db/otpauth"
	"time"
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
	otpAuthGetDB := db.NewOTPAuthGetDB(s.postgresConn)
	res, err := otpAuthGetDB.GetOTPAuth(ctx, db.GetOTPAuthParam{
		UserID: req.UserID,
		OTP:    req.OTP,
		Order: sqlorder.SQLOrder{
			Column: "created_at",
			By:     sqlorder.SQLOrderEnumASC,
		},
	})
	if err != nil {
		return ValidateOTPRes{}, err
	}

	if res.Status.String() != domain.OTPAuthStatusCreated {
		return ValidateOTPRes{}, db.ErrGetOTPAuthNotFound
	}

	otpAuthUpdateDB := db.NewOTPAuthUpdateDB(s.postgresConn)
	if time.Now().After(res.OTPExpiredAt) {
		err = otpAuthUpdateDB.UpdateOTPAuthStatus(ctx, db.UpdateOTPAuthStatusParam{ID: res.ID, Status: domain.OTPAuthStatusEnumExpired})
		if err != nil {
			// status expired is not automatically updated
			// logging
		}

		return ValidateOTPRes{}, db.ErrGetOTPAuthNotFound
	}

	err = otpAuthUpdateDB.UpdateOTPAuthStatus(ctx, db.UpdateOTPAuthStatusParam{ID: res.ID, Status: domain.OTPAuthStatusEnumValidated})
	if err != nil {
		// logging
	}

	return ValidateOTPRes{
		UserID: res.ID,
	}, nil
}
