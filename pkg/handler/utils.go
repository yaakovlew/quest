package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

func checkAuthHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != os.Getenv("AUTHORIZATION") {
			httpErrorResponse(w, http.StatusForbidden, fmt.Errorf("not valid header"))
			return
		}

		next.ServeHTTP(w, r)
	})

}
