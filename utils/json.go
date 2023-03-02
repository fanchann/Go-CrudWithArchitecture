package utils

import (
	"encoding/json"
	"net/http"
)

func ReadRequestFromBody(request *http.Request, form interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(form)
	ErrorWithPanic(err)
}

func WriteToResponse(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	ErrorWithPanic(err)
}
