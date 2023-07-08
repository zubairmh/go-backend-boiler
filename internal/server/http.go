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
	host := fmt.Sprintf("localhost:%d", port)
	log.Printf("Starting web server: %s", host)
	return &http.Server{
		Addr:    host,
		Handler: router,
	}
}
