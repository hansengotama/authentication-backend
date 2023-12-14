package validateotpauthservice

import (
	"context"
	"errors"
	"github.com/hansengotama/authentication-backend/internal/domain"
	sqlorder "github.com/hansengotama/authentication-backend/internal/lib/sql"
	"github.com/hansengotama/authentication-backend/internal/repository/db/getotpauthdb"
	"github.com/hansengotama/authentication-backend/internal/repository/db/updateotpstatusauth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func Test_ValidateOTPAuthService(t *testing.T) {
	dummyErr := errors.New("unexpected")

	testCases := []struct {
		name                string
		param               ValidateOTPAuthParam
		res                 ValidateOTPAuthRes
		mockGetOTPAuthDB    func() *getotpauthdb.MockGetOTPAuthDBInterface
		mockUpdateOTPAuthDB func() *updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface
		expectedErr         error
	}{
		{
			name: "when successfully validate otp auth",
			param: ValidateOTPAuthParam{
				UserID: 1,
				OTP:    11111,
			},
			res: ValidateOTPAuthRes{
				UserID: 1,
			},
			mockGetOTPAuthDB: func() *getotpauthdb.MockGetOTPAuthDBInterface {
				getOTPAuthDB := new(getotpauthdb.MockGetOTPAuthDBInterface)
				getOTPAuthDB.On("GetOTPAuth", mock.Anything, mock.Anything, getotpauthdb.GetOTPAuthParam{
					UserID: 1,
					OTP:    11111,
					Order: sqlorder.SQLOrder{
						Column: "created_at",
						By:     sqlorder.SQLOrderEnumASC,
					},
				}).Return(domain.OtpAuth{
					ID:           1,
					Status:       domain.OTPAuthStatusEnumCreated,
					OTPExpiredAt: time.Now().Add(5 * time.Minute),
				}, nil)

				return getOTPAuthDB
			},
			mockUpdateOTPAuthDB: func() *updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface {
				updateOTPAuthDB := new(updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface)
				updateOTPAuthDB.On("UpdateOTPAuthStatus", mock.Anything, mock.Anything, updateotpstatusauth.UpdateOTPAuthStatusParam{
					ID:     1,
					Status: domain.OTPAuthStatusEnumValidated,
				}).Return(nil)

				return updateOTPAuthDB
			},
		},
		{
			name: "when failed validate otp auth on call GetOTPAuth repo function",
			param: ValidateOTPAuthParam{
				UserID: 1,
				OTP:    11111,
			},
			res: ValidateOTPAuthRes{
				UserID: 1,
			},
			mockGetOTPAuthDB: func() *getotpauthdb.MockGetOTPAuthDBInterface {
				getOTPAuthDB := new(getotpauthdb.MockGetOTPAuthDBInterface)
				getOTPAuthDB.On("GetOTPAuth", mock.Anything, mock.Anything, getotpauthdb.GetOTPAuthParam{
					UserID: 1,
					OTP:    11111,
					Order: sqlorder.SQLOrder{
						Column: "created_at",
						By:     sqlorder.SQLOrderEnumASC,
					},
				}).Return(domain.OtpAuth{}, dummyErr)

				return getOTPAuthDB
			},
			mockUpdateOTPAuthDB: func() *updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface {
				updateOTPAuthDB := new(updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface)

				return updateOTPAuthDB
			},
			expectedErr: dummyErr,
		},
		{
			name: "when failed validate otp auth on status is not created",
			param: ValidateOTPAuthParam{
				UserID: 1,
				OTP:    11111,
			},
			res: ValidateOTPAuthRes{
				UserID: 1,
			},
			mockGetOTPAuthDB: func() *getotpauthdb.MockGetOTPAuthDBInterface {
				getOTPAuthDB := new(getotpauthdb.MockGetOTPAuthDBInterface)
				getOTPAuthDB.On("GetOTPAuth", mock.Anything, mock.Anything, getotpauthdb.GetOTPAuthParam{
					UserID: 1,
					OTP:    11111,
					Order: sqlorder.SQLOrder{
						Column: "created_at",
						By:     sqlorder.SQLOrderEnumASC,
					},
				}).Return(domain.OtpAuth{
					ID:           1,
					Status:       domain.OTPAuthStatusEnumExpired,
					OTPExpiredAt: time.Now().Add(5 * time.Minute),
				}, nil)

				return getOTPAuthDB
			},
			mockUpdateOTPAuthDB: func() *updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface {
				updateOTPAuthDB := new(updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface)

				return updateOTPAuthDB
			},
			expectedErr: getotpauthdb.ErrGetOTPAuthNotFound,
		},
		{
			name: "when failed validate otp auth on otp is already expired",
			param: ValidateOTPAuthParam{
				UserID: 1,
				OTP:    11111,
			},
			res: ValidateOTPAuthRes{
				UserID: 1,
			},
			mockGetOTPAuthDB: func() *getotpauthdb.MockGetOTPAuthDBInterface {
				getOTPAuthDB := new(getotpauthdb.MockGetOTPAuthDBInterface)
				getOTPAuthDB.On("GetOTPAuth", mock.Anything, mock.Anything, getotpauthdb.GetOTPAuthParam{
					UserID: 1,
					OTP:    11111,
					Order: sqlorder.SQLOrder{
						Column: "created_at",
						By:     sqlorder.SQLOrderEnumASC,
					},
				}).Return(domain.OtpAuth{
					ID:           1,
					Status:       domain.OTPAuthStatusEnumCreated,
					OTPExpiredAt: time.Now().Add(-5 * time.Minute),
				}, nil)

				return getOTPAuthDB
			},
			mockUpdateOTPAuthDB: func() *updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface {
				updateOTPAuthDB := new(updateotpstatusauth.MockUpdateOTPAuthStatusDBInterface)
				updateOTPAuthDB.On("UpdateOTPAuthStatus", mock.Anything, mock.Anything, updateotpstatusauth.UpdateOTPAuthStatusParam{
					ID:     1,
					Status: domain.OTPAuthStatusEnumExpired,
				}).Return(nil)

				return updateOTPAuthDB
			},
			expectedErr: getotpauthdb.ErrGetOTPAuthNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewValidateOTPAuthService(Dependency{
				GetOTPAuthDB:    tc.mockGetOTPAuthDB(),
				UpdateOTPAuthDB: tc.mockUpdateOTPAuthDB(),
			})

			res, err := s.ValidateOTPAuth(context.TODO(), tc.param)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedErr.Error())
				return
			}

			assert.Equal(t, tc.res.UserID, res.UserID)
			assert.NoError(t, err)
		})
	}
}
