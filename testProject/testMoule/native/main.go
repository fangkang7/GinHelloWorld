package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes map[string]http.HandlerFunc
}

func (r *Router) HandleFunc(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := r.routes[req.URL.Path]; ok {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func main() {
	router := &Router{routes: make(map[string]http.HandlerFunc)}

	router.HandleFunc("/kaka", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Println(123)
	})

	http.ListenAndServe(":8089", router)
}
