package main

import (
	"sync"

	"github.com/zubairmh/go-backend-boiler/internal/server"
)

func main() {
	var wg sync.WaitGroup
	instances := 3
	for i := 0; i < instances; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			server.CreateHTTPServer(5000 + i).ListenAndServe()
		}(i)
	}
	wg.Wait()
}
