package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateHTTPServer(port int) *http.Server {
	router := httprouter.New()
	createRoutes(router)
	log.Println("Starting web server")
	return &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: router,
	}
}
