package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateHTTPServer() *http.Server {
	router := httprouter.New()
	createRoutes(router)
	log.Println("Starting web server")
	return &http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}
}
