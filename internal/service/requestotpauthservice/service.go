package requestotpauthservice

import (
	"context"
	"github.com/hansengotama/authentication-backend/internal/domain"
	"github.com/hansengotama/authentication-backend/internal/lib/env"
	"github.com/hansengotama/authentication-backend/internal/lib/generator"
	"github.com/hansengotama/authentication-backend/internal/lib/postgres"
	"github.com/hansengotama/authentication-backend/internal/repository/db/insertotpauthdb"
	"time"
)

type RequestOTPAuthParam struct {
	UserID int
}

type RequestOTPAuthRes struct {
	UserID int
	OTP    int
}

type RequestOTPAuthServiceInterface interface {
	RequestOTPAuth(context.Context, RequestOTPAuthParam) (RequestOTPAuthRes, error)
}

type RequestOTPAuthService struct {
	Dependency Dependency
}

type Dependency struct {
	InsertOTPAuthDB insertotpauthdb.InsertOTPAuthDBInterface
}

func NewRequestOTPAuthService(dependency Dependency) RequestOTPAuthServiceInterface {
	return RequestOTPAuthService{
		Dependency: dependency,
	}
}

func (s RequestOTPAuthService) RequestOTPAuth(ctx context.Context, param RequestOTPAuthParam) (RequestOTPAuthRes, error) {
	otp, err := generator.RandomNumbers(5)
	if err != nil {
		return RequestOTPAuthRes{}, err
	}

	postgresConn := postgres.GetConnection()
	dbParam := insertotpauthdb.InsertOTPAuthParam{
		UserID:       param.UserID,
		OTP:          otp,
		OTPExpiredAt: time.Now().Add(env.GetOTPExpirationTime()),
		Status:       domain.OTPAuthStatusEnumCreated,
	}
	err = s.Dependency.InsertOTPAuthDB.InsertOTPAuth(ctx, postgresConn, dbParam)
	if err != nil {
		return RequestOTPAuthRes{}, err
	}

	return RequestOTPAuthRes{
		UserID: param.UserID,
		OTP:    otp,
	}, nil
}
