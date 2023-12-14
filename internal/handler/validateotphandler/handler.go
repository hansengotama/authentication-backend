package validateotphandler

import (
	"encoding/json"
	"errors"
	"github.com/hansengotama/authentication-backend/internal/lib/httphelper"
	"github.com/hansengotama/authentication-backend/internal/repository/db/getotpauthdb"
	"github.com/hansengotama/authentication-backend/internal/repository/db/updateotpstatusauth"
	"github.com/hansengotama/authentication-backend/internal/service/validateotpauthservice"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ValidateOTPAuthBody struct {
	UserID int `json:"user_id"`
	OTP    int `json:"otp"`
}

func (r ValidateOTPAuthBody) Validate() error {
	if r.UserID <= 0 {
		return errors.New("user id should be a positive integer")
	}

	if r.OTP <= 0 {
		return errors.New("otp should be a positive integer")
	}

	otpStr := strconv.Itoa(r.OTP)
	if len(otpStr) != 5 {
		return errors.New("otp must have 5 digits")
	}

	return nil
}

func (r ValidateOTPAuthBody) ToServiceParam() validateotpauthservice.ValidateOTPAuthParam {
	return validateotpauthservice.ValidateOTPAuthParam{
		UserID: r.UserID,
		OTP:    r.OTP,
	}
}

func HandleValidateOTPAuth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httphelper.Response(w, httphelper.HTTPResponse{
			Code:       http.StatusBadRequest,
			ErrMessage: httphelper.ErrReadingRequestBody.Error(),
		})
		return
	}

	var request ValidateOTPAuthBody
	err = json.Unmarshal(body, &request)
	if err != nil {
		httphelper.Response(w, httphelper.HTTPResponse{
			Code:       http.StatusBadRequest,
			ErrMessage: httphelper.ErrParsingRequestBody.Error(),
		})
		return
	}

	dep := validateotpauthservice.Dependency{
		GetOTPAuthDB:    getotpauthdb.GetOTPAuthDB{},
		UpdateOTPAuthDB: updateotpstatusauth.UpdateOTPAuthStatusDB{},
	}
	s := validateotpauthservice.NewValidateOTPAuthService(dep)
	res, err := s.ValidateOTPAuth(r.Context(), request.ToServiceParam())
	if err != nil {
		httphelper.Response(w, httphelper.HTTPResponse{
			Code:       http.StatusInternalServerError,
			ErrMessage: err.Error(),
		})
		return
	}

	httphelper.Response(w, httphelper.HTTPResponse{
		Data: res,
		Code: http.StatusOK,
	})
}
