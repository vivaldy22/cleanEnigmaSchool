package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	error2 "github.com/vivaldy22/cleanEnigmaSchool/tools/errors"
	"github.com/vivaldy22/cleanEnigmaSchool/tools/viper"
)

func InitDB() *sql.DB {
	dbUser := viper.ViperGetEnv("DB_USER", "root")
	dbPass := viper.ViperGetEnv("DB_PASSWORD", "password")
	dbHost := viper.ViperGetEnv("DB_HOST", "localhost")
	dbPort := viper.ViperGetEnv("DB_PORT", "3306")
	schemaName := viper.ViperGetEnv("DB_SCHEMA", "schema") // enigma_school_api

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	dbConn, _ := sql.Open("mysql", dbPath)
	err := dbConn.Ping()
	error2.PanicErr(err)
	return dbConn
}
