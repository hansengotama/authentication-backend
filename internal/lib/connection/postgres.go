package postgres

import (
	"database/sql"
	"fmt"
	"github.com/hansengotama/authentication-backend/internal/lib/env"
	_ "github.com/lib/pq"
)

var dbConn *sql.DB

func init() {
	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.GetPostgresDBUser(),
		env.GetPostgresDBPassword(),
		env.GetPostgresDBHost(),
		env.GetPostgresDBPort(),
		env.GetPostgresDBName(),
	)
	postgresDBCon, err := sql.Open("postgres", connectString)
	if err != nil {
		panic(err)
	}

	dbConn = postgresDBCon
}

func GetConnection() *sql.DB {
	return dbConn
}

func CloseConnection() {
	dbConn.Close()
}
