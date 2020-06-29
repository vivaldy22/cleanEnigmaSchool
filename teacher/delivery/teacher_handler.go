package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
	"github.com/vivaldy22/cleanEnigmaSchool/tools"
	"log"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TeacherHandler struct {
	TUseCase models.TeacherUseCase
}

func NewTeacherHandler(tu models.TeacherUseCase, router *mux.Router) {
	handler := &TeacherHandler{TUseCase: tu}
	router.HandleFunc("/teachers", handler.FetchTeachers).Methods("GET")
	router.HandleFunc("/teacher", handler.GetTeacherByID).Methods("GET")
	router.HandleFunc("/teacher", handler.InsertTeacher).Methods("POST")
	router.HandleFunc("/teacher", handler.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/teacher", handler.DeleteTeacher).Methods("DELETE")
}

func (t *TeacherHandler) FetchTeachers(w http.ResponseWriter, r *http.Request) {
	rawData, err := t.TUseCase.Fetch()
	tools.PrintlnErr(err)
	//var resp = response.Response{Msg: "Data Teacher", Data: getAll(db)}
	data, err := json.Marshal(rawData)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(data)
		log.Println("Endpoint hit: FetchTeachers")
	}
}

func (t *TeacherHandler) GetTeacherByID(w http.ResponseWriter, r *http.Request) {
	id := tools.ReadQueryParam("id", r)
	rawData, err := t.TUseCase.GetByID(id)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		// resp := response.Response{Msg: "Data Teacher By ID", Data: getByID(db, id)}
		data, err := json.Marshal(rawData)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Error occurred"))
		} else {
			w.Header().Set("content-type", "application/json")
			w.Write(data)
			log.Println("Endpoint hit: GetTeacherByID")
		}
	}
}

func (t *TeacherHandler) InsertTeacher(w http.ResponseWriter, r *http.Request) {
	var ts []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&ts)
	tools.PrintlnErr(err)
	for _, tt := range ts {
		err := t.TUseCase.Store(tt)
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

func (t *TeacherHandler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	err := t.TUseCase.Delete(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		log.Println("Delete successful")
		w.Write([]byte("Delete successful"))
	}
}

func (t *TeacherHandler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	var tt models.Teacher
	err := json.NewDecoder(r.Body).Decode(&tt)
	tools.PrintlnErr(err)
	err = t.TUseCase.Update(tt)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		log.Println("Update successful")
		w.Write([]byte("Update successful"))
	}
}
