package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code int
	Msg  string
}

type ResponsePayload struct {
	Code int
	Data []byte
}

func ConvertToJson(payload any) ([]byte, error) {
	return json.Marshal(payload)
}

func JsonResponse(w http.ResponseWriter, json []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}
