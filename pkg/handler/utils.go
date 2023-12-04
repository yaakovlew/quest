package handler

import (
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"error"`
}

func httpErrorResponse(w http.ResponseWriter, code int, err error) {
	log.Println(fmt.Sprintf("error: %s", err.Error()))

	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}
