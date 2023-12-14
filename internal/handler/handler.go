package handler

import (
	"encoding/json"
	postgres "github.com/hansengotama/authentication-backend/internal/lib/connection"
	"github.com/hansengotama/authentication-backend/internal/service/otpauth"
	"io/ioutil"
	"net/http"
)

type HTTPResponse struct {
	Data       any    `json:"data"`
	Code       int    `json:"code"`
	ErrMessage string `json:"err_message"`
}

func writeResponse(w http.ResponseWriter, resp HTTPResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(resp.Code)

	byt, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(byt)
	if err != nil {
		panic(err)
	}
}

func RequestOTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var request otpauth.RequestOTPReq
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	s := otpauth.NewAuthRequestService(postgres.GetConnection())
	res, err := s.Request(r.Context(), request)
	if err != nil {
		http.Error(w, "Error on request otp", http.StatusBadRequest)
		return
	}

	writeResponse(w, HTTPResponse{
		Data: res,
		Code: http.StatusOK,
	})
}

func ValidateOTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var request otpauth.ValidateOTPReq
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	s := otpauth.NewAuthValidateService(postgres.GetConnection())
	res, err := s.Validate(r.Context(), request)
	if err != nil {
		http.Error(w, "Error on validate otp", http.StatusBadRequest)
		return
	}

	writeResponse(w, HTTPResponse{
		Data: res,
		Code: http.StatusOK,
	})
}
