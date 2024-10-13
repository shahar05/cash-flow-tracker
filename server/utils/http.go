package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSON200 sets the Content-Type header and encodes the response as JSON with a 200 OK status
func WriteJSONOk(w http.ResponseWriter, data interface{}) {
	response := ResponseStructure{
		Data:   data,
		Status: true,
		Error:  nil,
	}
	writeJSONResponse(w, http.StatusOK, response)
}

// WriteJSONError sets the Content-Type header and encodes an error response as JSON with the specified status code
func WriteJSONError(w http.ResponseWriter, errorCode int, errorMessage string) {
	if errorCode < 400 || errorCode > 599 {
		log.Printf("WriteJSONError: Invalid errorCode: %d", errorCode)
		errorCode = http.StatusInternalServerError // Default to 500 if out of range
	}
	response := ResponseStructure{
		Data:   nil,
		Status: false,
		Error: &ErrorInfo{
			Message: errorMessage,
			Code:    errorCode,
		},
	}
	writeJSONResponse(w, errorCode, response)
}

// WriteJSONResponse sets the Content-Type header, status code, and encodes the response as JSON
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("writeJSONResponse: Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// // WriteJSON200 sets the Content-Type header and encodes the response as JSON with a 200 OK status
// func WriteJSON200(w http.ResponseWriter, data interface{}) {
// 	writeJSONResponse(w, http.StatusOK, data)
// }

// // WriteJSONResponse sets the Content-Type header, status code, and encodes the response as JSON
// func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	if err := json.NewEncoder(w).Encode(data); err != nil {
// 		log.Printf("WriteJSONResponse: Error encoding response: %v", err)
// 		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 	}
// }
