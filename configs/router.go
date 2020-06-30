package configs

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_stdeliv "github.com/vivaldy22/cleanEnigmaSchool/master/student/delivery"
	_strepo "github.com/vivaldy22/cleanEnigmaSchool/master/student/repository"
	_stusecase "github.com/vivaldy22/cleanEnigmaSchool/master/student/usecase"
	_sudeliv "github.com/vivaldy22/cleanEnigmaSchool/master/subject/delivery"
	_surepo "github.com/vivaldy22/cleanEnigmaSchool/master/subject/repository"
	_suusecase "github.com/vivaldy22/cleanEnigmaSchool/master/subject/usecase"
	_tdeliv "github.com/vivaldy22/cleanEnigmaSchool/master/teacher/delivery"
	_trepo "github.com/vivaldy22/cleanEnigmaSchool/master/teacher/repository"
	_tusecase "github.com/vivaldy22/cleanEnigmaSchool/master/teacher/usecase"
)

func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

func RunServer(r *mux.Router) {
	host := "localhost"
	port := "3000"
	fmt.Println("Starting Web Server at port:", port)
	err := http.ListenAndServe(host+": "+port, r)
	if err != nil {
		log.Fatal(err)
	}
}

func InitRouters(db *sql.DB, router *mux.Router) {
	tRep := _trepo.NewTeacherRepo(db)
	tUsc := _tusecase.NewTeacherUseCase(tRep)
	_tdeliv.NewTeacherHandler(tUsc, router)

	stRep := _strepo.NewStudentRepo(db)
	stUsc := _stusecase.NewStudentUseCase(stRep)
	_stdeliv.NewStudentHandler(stUsc, router)

	suRep := _surepo.NewSubjectRepo(db)
	suUsc := _suusecase.NewSubjectUseCase(suRep)
	_sudeliv.NewSubjectHandler(suUsc, router)
}
