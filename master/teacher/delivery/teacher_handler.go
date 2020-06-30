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

type TeacherHandler struct {
	TUseCase models.TeacherUseCase
}

func NewTeacherHandler(tu models.TeacherUseCase, router *mux.Router) {
	handler := &TeacherHandler{TUseCase: tu}
	router.HandleFunc("/teachers", handler.ShowTeachers).Methods(http.MethodGet)
	router.HandleFunc("/teacher", handler.CreateTeacher).Methods(http.MethodPost)
	router.HandleFunc("/teacher/{id}", handler.GetTeacherByID).Methods(http.MethodGet)
	router.HandleFunc("/teacher/{id}", handler.UpdateTeacher).Methods(http.MethodPut)
	router.HandleFunc("/teacher/{id}", handler.RemoveTeacher).Methods(http.MethodDelete)
}

func (t *TeacherHandler) ShowTeachers(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	data, err := t.TUseCase.Fetch()
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("ShowTeachers Failed", http.StatusNotFound, err.Error())
	} else {
		log.Println("Endpoint hit: FetchTeachers")
		resp = msgJson.Response("Teachers Data", http.StatusOK, data)
	}
	msgJson.WriteJSON(resp, w)
}

func (t *TeacherHandler) GetTeacherByID(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	id := varMux.GetVarsMux("id", r)
	data, err := t.TUseCase.GetByID(id)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("GetTeacherByID Failed", http.StatusNotFound, err.Error())
	} else {
		log.Println("Endpoint hit: GetTeacherByID")
		resp = msgJson.Response("Teacher Data", http.StatusOK, data)
	}
	msgJson.WriteJSON(resp, w)
}

func (t *TeacherHandler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	var teacher *models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("Decode failed", http.StatusBadRequest, err.Error())
	} else {
		err = t.TUseCase.Store(teacher)
		if err != nil {
			log.Println(err)
			resp = msgJson.Response("CreateTeacher failed", http.StatusBadRequest, err.Error())
		} else {
			log.Println("Endpoint hit: CreateTeacher")
			resp = msgJson.Response("CreateTeacher success", http.StatusCreated, "Insert success")
		}
	}
	msgJson.WriteJSON(resp, w)
}

func (t *TeacherHandler) RemoveTeacher(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	id := varMux.GetVarsMux("id", r)
	err := t.TUseCase.Delete(id)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("Delete failed", http.StatusNotFound, err.Error())
	} else {
		log.Println("Endpoint hit: RemoveTeacher")
		resp = msgJson.Response("Delete success", http.StatusOK, "Delete success")
	}
	msgJson.WriteJSON(resp, w)
}

func (t *TeacherHandler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	var resp *msgJson.ResponseMessage
	var teacher *models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		log.Println(err)
		resp = msgJson.Response("Decode failed", http.StatusBadRequest, err.Error())
	} else {
		id := varMux.GetVarsMux("id", r)
		err = t.TUseCase.Update(id, teacher)
		if err != nil {
			log.Println(err)
			resp = msgJson.Response("Update failed", http.StatusNotFound, err.Error())
		} else {
			log.Println("Endpoint hit: UpdateTeacher")
			resp = msgJson.Response("Update success", http.StatusOK, "Update success")
		}
	}
	msgJson.WriteJSON(resp, w)
}
