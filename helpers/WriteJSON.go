package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorJson struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "Application/Json")
	w.WriteHeader(status)
	w.Write(dataByte)
}
