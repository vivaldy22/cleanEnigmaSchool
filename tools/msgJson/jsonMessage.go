package msgJson

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseMessage struct {
	Status  string
	Message string
	Result  interface{}
}

func Response(m string, c int, r interface{}, err error) *ResponseMessage {
	if err != nil {
		log.Println(err)
	}
	return &ResponseMessage{http.StatusText(c), m, r}
}

func WriteJSON(resp *ResponseMessage, w http.ResponseWriter) {
	jsonMsg, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMsg)
}
