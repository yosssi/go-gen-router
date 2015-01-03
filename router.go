// generated by gen-router; DO NOT EDIT
package main

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", topHandler).Methods("GET")
	r.HandleFunc("/users", usersIndexHandler).Methods("GET")
	r.HandleFunc("/users/{id}", usersShowHandler).Methods("GET")
	r.HandleFunc("/users", usersCreateHandler).Methods("POST")
	return r
}
