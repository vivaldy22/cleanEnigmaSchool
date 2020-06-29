// Vivaldy Andhira Suwandhi - Challenge 6.3 Enigma School API #3

package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_tdeliv"github.com/vivaldy22/cleanEnigmaSchool/teacher/delivery"
	_trepo "github.com/vivaldy22/cleanEnigmaSchool/teacher/repository"
	_tusecase "github.com/vivaldy22/cleanEnigmaSchool/teacher/usecase"
	"github.com/vivaldy22/cleanEnigmaSchool/tools"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := tools.ViperGetEnv("DB_USER", "root")
	dbPass := tools.ViperGetEnv("DB_PASSWORD", "password")
	dbHost := tools.ViperGetEnv("DB_HOST", "localhost")
	dbPort := tools.ViperGetEnv("DB_PORT", "3306")
	schemaName := tools.ViperGetEnv("DB_SCHEMA", "schema") // enigma_school_api

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	dbConn, _ := sql.Open("mysql", dbPath)
	err := dbConn.Ping()
	tools.PanicErr(err)
	defer func() {
		err := dbConn.Close()
		tools.FatalErr(err)
	}()

	router := mux.NewRouter()
	tr := _trepo.NewTeacherRepo(dbConn)
	tu := _tusecase.NewTeacherUseCase(tr)
	_tdeliv.NewTeacherHandler(tu, router)

	fmt.Println("Running on port 3000")
	err = http.ListenAndServe(":3000", nil)
	tools.FatalErr(err)
}
