package requestotphandler

import (
	"encoding/json"
	"errors"
	"github.com/hansengotama/authentication-backend/internal/lib/httphelper"
	"github.com/hansengotama/authentication-backend/internal/repository/db/insertotpauthdb"
	"github.com/hansengotama/authentication-backend/internal/service/requestotpauthservice"
	"io/ioutil"
	"net/http"
)

type RequestOTPAuthBody struct {
	UserID int `json:"user_id"`
}

func (r RequestOTPAuthBody) Validate() error {
	if r.UserID <= 0 {
		return errors.New("user id should be a positive integer")
	}

	return nil
}

func (r RequestOTPAuthBody) ToServiceParam() requestotpauthservice.RequestOTPAuthParam {
	return requestotpauthservice.RequestOTPAuthParam{
		UserID: r.UserID,
	}
}

func HandleRequestOTPAuth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httphelper.Response(w, httphelper.HTTPResponse{
			Code:       http.StatusBadRequest,
			ErrMessage: httphelper.ErrReadingRequestBody.Error(),
		})
		return
	}

	var request RequestOTPAuthBody
	err = json.Unmarshal(body, &request)
	if err != nil {
		httphelper.Response(w, httphelper.HTTPResponse{
			Code:       http.StatusBadRequest,
			ErrMessage: httphelper.ErrParsingRequestBody.Error(),
		})
		return
	}

	dep := requestotpauthservice.Dependency{
		InsertOTPAuthDB: insertotpauthdb.InsertOTPAuthDB{},
	}
	s := requestotpauthservice.NewRequestOTPAuthService(dep)
	res, err := s.RequestOTPAuth(r.Context(), request.ToServiceParam())
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
