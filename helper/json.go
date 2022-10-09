package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, response interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(response)
	PanicIfError(err)
}

func WriteToRequestBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
