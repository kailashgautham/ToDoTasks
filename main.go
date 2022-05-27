package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from To-Do Tasks"))
}

func showTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing todos..."))
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating todo..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/todo", showTodos)
	mux.HandleFunc("/todo/create", createTodo)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
