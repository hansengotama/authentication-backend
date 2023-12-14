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
	UserID int
	OTP    int
}

type ValidateOTPAuthRes struct {
	UserID int `json:"user_id"`
}

type ValidateOTPAuthServiceInterface interface {
	ValidateOTPAuth(context.Context, ValidateOTPAuthParam) (ValidateOTPAuthRes, error)
}

type ValidateOTPAuthService struct {
	Dependency Dependency
}

type Dependency struct {
	GetOTPAuthDB    getotpauthdb.GetOTPAuthDBInterface
	UpdateOTPAuthDB updateotpstatusauth.UpdateOTPAuthStatusDBInterface
}

func NewValidateOTPAuthService(dependency Dependency) ValidateOTPAuthServiceInterface {
	return ValidateOTPAuthService{
		Dependency: dependency,
	}
}

func (s ValidateOTPAuthService) ValidateOTPAuth(ctx context.Context, param ValidateOTPAuthParam) (ValidateOTPAuthRes, error) {
	postgresConn := postgres.GetConnection()
	res, err := s.Dependency.GetOTPAuthDB.GetOTPAuth(ctx, postgresConn, getotpauthdb.GetOTPAuthParam{
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

	isOTPExpired := time.Now().After(res.OTPExpiredAt)
	if isOTPExpired {
		err = s.Dependency.UpdateOTPAuthDB.UpdateOTPAuthStatus(
			ctx,
			postgresConn,
			updateotpstatusauth.UpdateOTPAuthStatusParam{
				ID:     res.ID,
				Status: domain.OTPAuthStatusEnumExpired,
			})
		if err != nil {
			// status expired is not automatically updated
			// logging
		}

		return ValidateOTPAuthRes{}, getotpauthdb.ErrGetOTPAuthNotFound
	}

	err = s.Dependency.UpdateOTPAuthDB.UpdateOTPAuthStatus(
		ctx,
		postgresConn,
		updateotpstatusauth.UpdateOTPAuthStatusParam{
			ID:     res.ID,
			Status: domain.OTPAuthStatusEnumValidated,
		})
	if err != nil {
		// logging
	}

	return ValidateOTPAuthRes{
		UserID: res.ID,
	}, nil
}
