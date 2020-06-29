// Vivaldy Andhira Suwandhi - Challenge 6.3 Enigma School API #3

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_stdeliv "github.com/vivaldy22/cleanEnigmaSchool/student/delivery"
	_strepo "github.com/vivaldy22/cleanEnigmaSchool/student/repository"
	_stusecase "github.com/vivaldy22/cleanEnigmaSchool/student/usecase"
	_sudeliv "github.com/vivaldy22/cleanEnigmaSchool/subject/delivery"
	_surepo "github.com/vivaldy22/cleanEnigmaSchool/subject/repository"
	_suusecase "github.com/vivaldy22/cleanEnigmaSchool/subject/usecase"
	_tdeliv "github.com/vivaldy22/cleanEnigmaSchool/teacher/delivery"
	_trepo "github.com/vivaldy22/cleanEnigmaSchool/teacher/repository"
	_tusecase "github.com/vivaldy22/cleanEnigmaSchool/teacher/usecase"
	"github.com/vivaldy22/cleanEnigmaSchool/tools"

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
	http.Handle("/", router)

	tRep := _trepo.NewTeacherRepo(dbConn)
	tUsc := _tusecase.NewTeacherUseCase(tRep)
	_tdeliv.NewTeacherHandler(tUsc, router)

	stRep := _strepo.NewStudentRepo(dbConn)
	stUsc := _stusecase.NewStudentUseCase(stRep)
	_stdeliv.NewStudentHandler(stUsc, router)

	suRep := _surepo.NewSubjectRepo(dbConn)
	suUsc := _suusecase.NewSubjectUseCase(suRep)
	_sudeliv.NewSubjectHandler(suUsc, router)

	fmt.Println("Running on port 3000")
	err = http.ListenAndServe(":3000", nil)
	tools.FatalErr(err)
}
