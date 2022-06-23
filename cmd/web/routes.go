package main

import (
	"net/http"
)

func (app *Application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/todo", app.showTodo)
	mux.HandleFunc("/todo/create", app.createTodo)
	mux.HandleFunc("/todo/today", app.showTodayTodos)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
