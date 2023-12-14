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

func GetPostgresDBHost() string {
	dbHost := os.Getenv("POSTGRES_DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	return dbHost
}

func GetPostgresDBPort() string {
	dbPort := os.Getenv("POSTGRES_DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	return dbPort
}

func GetPostgresDBName() string {
	dbName := os.Getenv("POSTGRES_DB_NAME")
	if dbName == "" {
		dbName = "authentication"
	}

	return dbName
}

func GetPostgresDBUser() string {
	dbUser := os.Getenv("POSTGRES_DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	return dbUser
}

func GetPostgresDBPassword() string {
	dbPassword := os.Getenv("POSTGRES_DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}

	return dbPassword
}

func GetOTPExpirationTime() time.Duration {
	expirationTime := os.Getenv("OTP_EXPIRATION_TIME")
	if expirationTime != "" {
		duration, err := time.ParseDuration(expirationTime)
		if err == nil {
			return duration
		}
	}

	defaultExpirationTime := 2 * time.Minute
	return defaultExpirationTime
}
