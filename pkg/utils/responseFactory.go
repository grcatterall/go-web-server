package utils

import (
	"fmt"
	"net/http"
)

type ResponseFactory struct{}

func NewResponseFactory() *ResponseFactory {
	return &ResponseFactory{}
}

func (rf *ResponseFactory) SuccessResponse(w http.ResponseWriter, status int, body []byte) {
	JsonResponse(w, body, status)
}

func (rf *ResponseFactory) ErrorResponse(w http.ResponseWriter, status int, message string) {
	jsonMsg, err := ConvertToJson(map[string]string{"error": message})

	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	rf.SuccessResponse(w, status, jsonMsg)
}

func (rf *ResponseFactory) ResponseDefer(w http.ResponseWriter) {
	if err := recover(); err != nil {
		if error, ok := err.(ErrorResponse); ok {
			fmt.Printf("Error message: %s\n", error.Msg)
			fmt.Printf("Error code: %d\n", error.Code)
			rf.ErrorResponse(w, error.Code, error.Msg)

		} else {
			fmt.Println("Recovered from panic:", err)
			rf.ErrorResponse(w, 500, "Server Error")
		}
	}
}
