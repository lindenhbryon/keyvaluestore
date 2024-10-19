package server

import (
	"fmt"
	"net/http"
)

func routes() {
	http.HandleFunc("/", mainHandler())
}

func Setup() {
	routes()
}

func mainHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")

		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "method get")
		case http.MethodPut:

		}
	}
}
