package otpauth

type OtpAuthVerify interface {
	verify() error
}
