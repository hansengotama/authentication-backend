package env

import (
	"os"
	"time"
)

func GetAppPort() string {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	return appPort
}

func GetOTPExpirationTime() time.Duration {
	expirationTime := os.Getenv("EXPIRATION_TIME")
	if expirationTime == "" {
		duration, err := time.ParseDuration(expirationTime)
		if err == nil {
			return duration
		}

		// logging.Err
	}

	defaultExpirationTime := time.Duration(2 * time.Minute)
	return defaultExpirationTime
}
