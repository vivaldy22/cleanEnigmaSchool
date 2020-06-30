package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
	_error "github.com/vivaldy22/cleanEnigmaSchool/tools/errors"
	"github.com/vivaldy22/cleanEnigmaSchool/tools/varMux"
)

type ResponseError struct {
	Message string `json:"message"`
}

type StudentHandler struct {
	StUseCase models.StudentUseCase
}

func NewStudentHandler(tu models.StudentUseCase, router *mux.Router) {
	handler := &StudentHandler{StUseCase: tu}
	router.HandleFunc("/students", handler.FetchStudents).Methods("GET")
	router.HandleFunc("/student", handler.GetStudentByID).Methods("GET")
	router.HandleFunc("/student", handler.InsertStudent).Methods("POST")
	router.HandleFunc("/student", handler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student", handler.DeleteStudent).Methods("DELETE")
}

func (s *StudentHandler) FetchStudents(w http.ResponseWriter, r *http.Request) {
	rawData, err := s.StUseCase.Fetch()
	_error.PrintlnErr(err)
	//var resp = response.Response{Msg: "Data Student", Data: getAll(db)}
	data, err := json.Marshal(rawData)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(data)
		log.Println("Endpoint hit: FetchStudents")
	}
}

func (s *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	rawData, err := s.StUseCase.GetByID(id)
	// resp := response.Response{Msg: "Data Student By ID", Data: getByID(db, id)}
	data, err := json.Marshal(rawData)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(data)
		log.Println("Endpoint hit: GetStudentByID")
	}
}

func (s *StudentHandler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	var sts []models.Student
	err := json.NewDecoder(r.Body).Decode(&sts)
	_error.PrintlnErr(err)
	for _, ss := range sts {
		err := s.StUseCase.Store(ss)
		_error.PrintlnErr(err)
	}
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		log.Println("Insert successful")
		w.Write([]byte("Insert successful"))
	}
}

func (s *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	err := s.StUseCase.Delete(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		log.Println("Delete successful")
		w.Write([]byte("Delete successful"))
	}
}

func (s *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var ss models.Student
	err := json.NewDecoder(r.Body).Decode(&ss)
	_error.PrintlnErr(err)
	err = s.StUseCase.Update(ss)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		log.Println("Update successful")
		w.Write([]byte("Update successful"))
	}
}
