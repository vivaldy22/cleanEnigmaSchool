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

type SubjectHandler struct {
	SUseCase models.SubjectUseCase
}

func NewSubjectHandler(tu models.SubjectUseCase, router *mux.Router) {
	handler := &SubjectHandler{SUseCase: tu}
	router.HandleFunc("/subjects", handler.ShowSubjects).Methods(http.MethodGet)
	router.HandleFunc("/subject", handler.CreateSubject).Methods(http.MethodPost)
	router.HandleFunc("/subject/{id}", handler.GetSubjectByID).Methods(http.MethodGet)
	router.HandleFunc("/subject/{id}", handler.UpdateSubject).Methods(http.MethodPut)
	router.HandleFunc("/subject/{id}", handler.RemoveSubject).Methods(http.MethodDelete)
}

func (s *SubjectHandler) ShowSubjects(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	data, err := s.SUseCase.Fetch()
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("ShowSubjects Failed", http.StatusNotFound, err.Error())
	} else {
		log.Println("Endpoint hit: ShowSubjects")
		resp = msgJson.Response("Subjects Data", http.StatusOK, data)
	}
	msgJson.WriteJSON(resp, w)
}

func (s *SubjectHandler) GetSubjectByID(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	id := varMux.GetVarsMux("id", r)
	data, err := s.SUseCase.GetByID(id)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("GetSubjectByID Failed", http.StatusNotFound, err.Error())
	} else {
		log.Println("Endpoint hit: GetSubjectByID")
		resp = msgJson.Response("Subject Data", http.StatusOK, data)
	}
	msgJson.WriteJSON(resp, w)
}

func (s *SubjectHandler) CreateSubject(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	var subject *models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("Decode failed", http.StatusBadRequest, err.Error())
	} else {
		err = s.SUseCase.Store(subject)
		if err != nil {
			log.Println(err)
			resp = msgJson.Response("CreateSubject failed", http.StatusBadRequest, err.Error())
		} else {
			log.Println("Endpoint hit: CreateSubject")
			resp = msgJson.Response("CreateSubject success", http.StatusCreated, "Insert success")
		}
	}
	msgJson.WriteJSON(resp, w)
}

func (s *SubjectHandler) RemoveSubject(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	id := varMux.GetVarsMux("id", r)
	if _, err := s.SUseCase.GetByID(id); err != nil {
		log.Println(err)
		resp = msgJson.Response("Data not found", http.StatusNotFound, err.Error())
	} else {
		err := s.SUseCase.Delete(id)
		if err != nil {
			log.Println(err)
			resp = msgJson.Response("Delete failed", http.StatusNotFound, err.Error())
		} else {
			log.Println("Endpoint hit: RemoveSubject")
			resp = msgJson.Response("Delete success", http.StatusOK, "Delete success")
		}
	}
	msgJson.WriteJSON(resp, w)
}

func (s *SubjectHandler) UpdateSubject(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	var subject *models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("Decode failed", http.StatusBadRequest, err.Error())
	} else {
		id := varMux.GetVarsMux("id", r)
		_, err := s.SUseCase.GetByID(id)
		if err != nil {
			log.Println(err)
			resp = msgJson.Response("Data not found", http.StatusNotFound, err.Error())
		} else {
			err = s.SUseCase.Update(id, subject)
			if err != nil {
				log.Println(err)
				resp = msgJson.Response("Update failed", http.StatusNotFound, err.Error())
			} else {
				log.Println("Endpoint hit: UpdateSubject")
				resp = msgJson.Response("Update success", http.StatusOK, "Update success")
			}
		}
	}
	msgJson.WriteJSON(resp, w)
}
