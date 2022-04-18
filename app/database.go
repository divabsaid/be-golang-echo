package app

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func InitDatabase(config *viper.Viper) *sql.DB {
	// init database
	dbHost := config.GetString("DB_HOST")
	dbPort := config.GetString(`DB_PORT`)
	dbUser := config.GetString(`DB_USER`)
	dbPass := config.GetString(`DB_PASS`)
	dbName := config.GetString(`DB_NAME`)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		panic(fmt.Errorf("fatal error database connection: %s", err))
	}
	return dbConn
}
