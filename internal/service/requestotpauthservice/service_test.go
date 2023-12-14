package requestotpauthservice

import (
	"context"
	"errors"
	"github.com/hansengotama/authentication-backend/internal/repository/db/insertotpauthdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strconv"
	"testing"
)

func Test_RequestOTPAuthService(t *testing.T) {
	dummyErr := errors.New("unexpected")

	testCases := []struct {
		name                string
		param               RequestOTPAuthParam
		res                 RequestOTPAuthRes
		mockInsertOTPAuthDB func() *insertotpauthdb.MockInsertOTPAuthDBInterface
		expectedErr         error
	}{
		{
			name: "when successfully request otp auth",
			param: RequestOTPAuthParam{
				UserID: 1,
			},
			res: RequestOTPAuthRes{
				UserID: 1,
			},
			mockInsertOTPAuthDB: func() *insertotpauthdb.MockInsertOTPAuthDBInterface {
				insertOTPAuthDB := new(insertotpauthdb.MockInsertOTPAuthDBInterface)
				insertOTPAuthDB.On("InsertOTPAuth", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				return insertOTPAuthDB
			},
		},
		{
			name: "when failed request otp auth on call InsertOTPAuth repo function",
			param: RequestOTPAuthParam{
				UserID: 1,
			},
			res: RequestOTPAuthRes{
				UserID: 1,
			},
			mockInsertOTPAuthDB: func() *insertotpauthdb.MockInsertOTPAuthDBInterface {
				insertOTPAuthDB := new(insertotpauthdb.MockInsertOTPAuthDBInterface)
				insertOTPAuthDB.On("InsertOTPAuth", mock.Anything, mock.Anything, mock.Anything).Return(dummyErr)

				return insertOTPAuthDB
			},
			expectedErr: dummyErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewRequestOTPAuthService(Dependency{
				InsertOTPAuthDB: tc.mockInsertOTPAuthDB(),
			})

			res, err := s.RequestOTPAuth(context.TODO(), tc.param)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedErr.Error())
				return
			}

			assert.Equal(t, tc.res.UserID, res.UserID)

			otpStr := strconv.Itoa(res.OTP)
			assert.Equal(t, 5, len(otpStr))
			assert.NoError(t, err)
		})
	}
}
