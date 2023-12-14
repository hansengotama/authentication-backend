package httphelper

import (
	"encoding/json"
	"errors"
	"net/http"
)

type HTTPResponse struct {
	Data       any    `json:"data"`
	Code       int    `json:"code"`
	ErrMessage string `json:"err_message"`
}

func Response(w http.ResponseWriter, resp HTTPResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(resp.Code)

	byt, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(byt)
	if err != nil {
		panic(err)
	}
}

var ErrReadingRequestBody = errors.New("error reading request body")
var ErrParsingRequestBody = errors.New("error parsing request body")
