package utils

import (
	"net/http"
    "encoding/json"
)


func ConvertToJson(payload any) ([]byte, error) {
    return json.Marshal(payload)
}

func JsonResponse(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(json)
}