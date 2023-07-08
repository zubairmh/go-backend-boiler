package main

import (
	"github.com/zubairmh/go-backend-boiler/internal/server"
)

func main() {
	instances := 3
	for i := 0; i < instances; i++ {
		go server.CreateHTTPServer(5000 + instances).ListenAndServe()
	}
	for {
	}
}
