package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
