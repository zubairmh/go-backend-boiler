package main

import (
	"github.com/zubairmh/go-backend-boiler/internal/server"
)

func main() {
	server.CreateHTTPServer().ListenAndServe()
}
