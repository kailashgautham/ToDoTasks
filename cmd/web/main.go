package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/todo", showTodo)
	mux.HandleFunc("/todo/create", createTodo)
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static", fileServer)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
