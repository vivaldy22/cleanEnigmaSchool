package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
	"github.com/vivaldy22/cleanEnigmaSchool/tools/msgJson"
	"github.com/vivaldy22/cleanEnigmaSchool/tools/varMux"
)

type StudentHandler struct {
	StUseCase models.StudentUseCase
}

func NewStudentHandler(tu models.StudentUseCase, router *mux.Router) {
	handler := &StudentHandler{StUseCase: tu}
	router.HandleFunc("/students", handler.FetchStudents).Methods(http.MethodGet)
	router.HandleFunc("/student", handler.InsertStudent).Methods(http.MethodPost)
	router.HandleFunc("/student/{id}", handler.GetStudentByID).Methods(http.MethodGet)
	router.HandleFunc("/student/{id}", handler.UpdateStudent).Methods(http.MethodPut)
	router.HandleFunc("/student/{id}", handler.RemoveStudent).Methods(http.MethodDelete)
}

func (s *StudentHandler) FetchStudents(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	data, err := s.StUseCase.Fetch()
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("ShowStudents Failed", http.StatusNotFound, nil, err)
	} else {
		log.Println("Endpoint hit: FetchStudents")
		resp = msgJson.Response("Students Data", http.StatusOK, data, nil)
	}
	msgJson.WriteJSON(resp, w)
}

func (s *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	id := varMux.GetVarsMux("id", r)
	data, err := s.StUseCase.GetByID(id)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("GetStudentByID Failed", http.StatusNotFound, nil, err)
	} else {
		log.Println("Endpoint hit: GetStudentByID")
		resp = msgJson.Response("Student Data", http.StatusOK, data, nil)
	}
	msgJson.WriteJSON(resp, w)
}

func (s *StudentHandler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	var student *models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("Decode failed", http.StatusBadRequest, nil, err)
	} else {
		err = s.StUseCase.Store(student)
		if err != nil {
			log.Println(err)
			resp = msgJson.Response("CreateStudent failed", http.StatusBadRequest, nil, err)
		} else {
			log.Println("Endpoint hit: CreateStudent")
			resp = msgJson.Response("CreateStudent success", http.StatusCreated, student, nil)
		}
	}
	msgJson.WriteJSON(resp, w)
}

func (s *StudentHandler) RemoveStudent(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	id := varMux.GetVarsMux("id", r)
	data, err := s.StUseCase.GetByID(id)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("Data not found", http.StatusNotFound, nil, err)
	} else {
		err := s.StUseCase.Delete(id)
		if err != nil {
			log.Println(err)
			resp = msgJson.Response("Delete failed", http.StatusNotFound, nil, err)
		} else {
			log.Println("Endpoint hit: RemoveStudent")
			resp = msgJson.Response("Delete success", http.StatusOK, data, nil)
		}
	}
	msgJson.WriteJSON(resp, w)
}

func (s *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	var student *models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("Decode failed", http.StatusBadRequest, nil, err)
	} else {
		id := varMux.GetVarsMux("id", r)
		data, err := s.StUseCase.GetByID(id)
		if err != nil {
			log.Println(err)
			resp = msgJson.Response("Data not found", http.StatusNotFound, nil, err)
		} else {
			err = s.StUseCase.Update(id, student)
			if err != nil {
				log.Println(err)
				resp = msgJson.Response("Update failed", http.StatusNotFound, nil, err)
			} else {
				log.Println("Endpoint hit: UpdateStudent")
				resp = msgJson.Response("Update success", http.StatusOK, data, nil)
			}
		}
	}
	msgJson.WriteJSON(resp, w)
}
