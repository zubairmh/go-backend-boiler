package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateHTTPServer() *http.Server {
	router := httprouter.New()
	return &http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}
}
