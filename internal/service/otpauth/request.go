package otpauth

import "context"

type RequestOTPReq struct {
	UserID int
}

type RequestOTPQueryRes struct {
	UserID int
	OTP    int
}

type OtpAuthRequestServiceInterface interface {
	Request(context.Context, RequestOTPReq) (RequestOTPQueryRes, error)
}

type OtpAuthRequestService struct{}

func NewAuthRequestService() OtpAuthRequestServiceInterface {
	return OtpAuthRequestService{}
}

func (s OtpAuthRequestService) Request(context.Context, RequestOTPReq) (RequestOTPQueryRes, error) {
	return RequestOTPQueryRes{}, nil
}
