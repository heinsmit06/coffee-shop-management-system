package utils

import (
	"encoding/json"
	"net/http"
)

type myErrorType struct {
	ErrorMessage string `json:"error"`
}

func ResponseErrorJson(err error, w http.ResponseWriter) {
	var errStruct myErrorType
	var statusCode int = 400
	errStruct.ErrorMessage = err.Error()

	// specify appropriate statusCode

	jsonData, _ := json.MarshalIndent(errStruct, "", " ")
	http.Error(w, string(jsonData), statusCode)
}
