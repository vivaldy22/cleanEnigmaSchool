package config

import (
	"database/sql"
	"fmt"
	"github.com/vivaldy22/cleanEnigmaSchool/tools"
)

func connectDB(dbPath string) (*sql.DB, error) {
	db, _ := sql.Open("mysql", dbPath)
	err := db.Ping()
	tools.PanicErr(err)
	return db, nil
}

func InitDB() *sql.DB {
	dbUser := tools.ViperGetEnv("DB_USER", "root")
	dbPass := tools.ViperGetEnv("DB_PASSWORD", "password")
	dbHost := tools.ViperGetEnv("DB_HOST", "localhost")
	dbPort := tools.ViperGetEnv("DB_PORT", "3306")
	schemaName := tools.ViperGetEnv("DB_SCHEMA", "schema") // enigma_school_api

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	db, err := connectDB(dbPath)
	tools.PanicErr(err)
	return db
}
