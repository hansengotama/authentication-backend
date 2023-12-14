package validateotpauthservice

import (
	"context"
	"github.com/hansengotama/authentication-backend/internal/domain"
	"github.com/hansengotama/authentication-backend/internal/lib/postgres"
	sqlorder "github.com/hansengotama/authentication-backend/internal/lib/sql"
	"github.com/hansengotama/authentication-backend/internal/repository/db/getotpauthdb"
	"github.com/hansengotama/authentication-backend/internal/repository/db/updateotpstatusauth"
	"time"
)

type ValidateOTPAuthParam struct {
	UserID int `json:"user_id"`
	OTP    int `json:"otp"`
}

type ValidateOTPAuthRes struct {
	UserID int `json:"user_id"`
}

type ValidateOTPAuthServiceInterface interface {
	ValidateOTPAuthService(context.Context, ValidateOTPAuthParam) (ValidateOTPAuthRes, error)
}

type ValidateOTPAuthService struct {
	dependency Dependency
}

type Dependency struct {
	GetOTPAuthDB    getotpauthdb.GetOTPAuthDB
	UpdateOTPAuthDB updateotpstatusauth.UpdateOTPAuthStatusDB
}

func NewValidateOTPAuthService(dependency Dependency) ValidateOTPAuthServiceInterface {
	return ValidateOTPAuthService{
		dependency: dependency,
	}
}

func (s ValidateOTPAuthService) ValidateOTPAuthService(ctx context.Context, param ValidateOTPAuthParam) (ValidateOTPAuthRes, error) {
	postgresConn := postgres.GetConnection()
	res, err := s.dependency.GetOTPAuthDB.GetOTPAuth(ctx, postgresConn, getotpauthdb.GetOTPAuthParam{
		UserID: param.UserID,
		OTP:    param.OTP,
		Order: sqlorder.SQLOrder{
			Column: "created_at",
			By:     sqlorder.SQLOrderEnumASC,
		},
	})
	if err != nil {
		return ValidateOTPAuthRes{}, err
	}

	if res.Status.String() != domain.OTPAuthStatusCreated {
		return ValidateOTPAuthRes{}, getotpauthdb.ErrGetOTPAuthNotFound
	}

	if time.Now().After(res.OTPExpiredAt) {
		err = s.dependency.UpdateOTPAuthDB.UpdateOTPAuthStatus(ctx, postgresConn, updateotpstatusauth.UpdateOTPAuthStatusParam{ID: res.ID, Status: domain.OTPAuthStatusEnumExpired})
		if err != nil {
			// status expired is not automatically updated
			// logging
		}

		return ValidateOTPAuthRes{}, getotpauthdb.ErrGetOTPAuthNotFound
	}

	err = s.dependency.UpdateOTPAuthDB.UpdateOTPAuthStatus(ctx, postgresConn, updateotpstatusauth.UpdateOTPAuthStatusParam{ID: res.ID, Status: domain.OTPAuthStatusEnumValidated})
	if err != nil {
		// logging
	}

	return ValidateOTPAuthRes{
		UserID: res.ID,
	}, nil
}
