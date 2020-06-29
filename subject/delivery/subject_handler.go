package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
	"github.com/vivaldy22/cleanEnigmaSchool/tools"
)

type ResponseError struct {
	Message string `json:"message"`
}

type SubjectHandler struct {
	SUseCase models.SubjectUseCase
}

func NewSubjectHandler(tu models.SubjectUseCase, router *mux.Router) {
	handler := &SubjectHandler{SUseCase: tu}
	router.HandleFunc("/subjects", handler.FetchSubjects).Methods("GET")
	router.HandleFunc("/subject", handler.GetSubjectByID).Methods("GET")
	router.HandleFunc("/subject", handler.InsertSubject).Methods("POST")
	router.HandleFunc("/subject", handler.UpdateSubject).Methods("PUT")
	router.HandleFunc("/subject", handler.DeleteSubject).Methods("DELETE")
}

func (s *SubjectHandler) FetchSubjects(w http.ResponseWriter, r *http.Request) {
	rawData, err := s.SUseCase.Fetch()
	tools.PrintlnErr(err)
	//var resp = response.Response{Msg: "Data Subject", Data: getAll(db)}
	data, err := json.Marshal(rawData)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(data)
		log.Println("Endpoint hit: FetchSubjects")
	}
}

func (s *SubjectHandler) GetSubjectByID(w http.ResponseWriter, r *http.Request) {
	id := tools.ReadQueryParam("id", r)
	rawData, err := s.SUseCase.GetByID(id)
	// resp := response.Response{Msg: "Data Subject By ID", Data: getByID(db, id)}
	data, err := json.Marshal(rawData)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(data)
		log.Println("Endpoint hit: GetSubjectByID")
	}
}

func (s *SubjectHandler) InsertSubject(w http.ResponseWriter, r *http.Request) {
	var sus []models.Subject
	err := json.NewDecoder(r.Body).Decode(&sus)
	tools.PrintlnErr(err)
	for _, ss := range sus {
		err := s.SUseCase.Store(ss)
		tools.PrintlnErr(err)
	}
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		log.Println("Insert successful")
		w.Write([]byte("Insert successful"))
	}
}

func (s *SubjectHandler) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	err := s.SUseCase.Delete(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		log.Println("Delete successful")
		w.Write([]byte("Delete successful"))
	}
}

func (s *SubjectHandler) UpdateSubject(w http.ResponseWriter, r *http.Request) {
	var sus models.Subject
	err := json.NewDecoder(r.Body).Decode(&sus)
	tools.PrintlnErr(err)
	err = s.SUseCase.Update(sus)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		log.Println("Update successful")
		w.Write([]byte("Update successful"))
	}
}
