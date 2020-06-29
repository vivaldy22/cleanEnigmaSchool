// Vivaldy Andhira Suwandhi - Challenge 6.3 Enigma School API #3

package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vivaldy22/cleanEnigmaSchool/delivery"
	repo "github.com/vivaldy22/cleanEnigmaSchool/repositories"
	"github.com/vivaldy22/cleanEnigmaSchool/tools"
	"github.com/vivaldy22/cleanEnigmaSchool/usecase"
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
	tr := repo.NewTeacherRepo(dbConn)
	tu := usecase.NewTeacherUseCase(tr)
	delivery.NewTeacherHandler(tu, router)

	//http.HandleFunc("/", index)
	//http.HandleFunc("/teachers", ac.GetTeacher())
	//http.HandleFunc("/teacher", ac.TeacherRouter())
	//http.HandleFunc("/students", ac.GetStudent())
	//http.HandleFunc("/student", ac.StudentRouter())
	//http.HandleFunc("/subjects", ac.GetSubject())
	//http.HandleFunc("/subject", ac.SubjectRouter())

	fmt.Println("Running on port 3000")
	err = http.ListenAndServe(":3000", nil)
	tools.FatalErr(err)
}
