package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"kailashgautham.com/todotasks/pkg/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	t, err := app.todos.ShowPending()

	if err != nil {
		app.serverError(w, err)
	}

	td := &templateData{Todos: t}

	app.render(w, r, "home.page.tmpl", td)

}

func (app *Application) showTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	todo, err := app.todos.OpenDetails(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	td := &templateData{Todo: todo}
	app.render(w, r, "show.page.tmpl", td)
}

func (app *Application) showTodayTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Showing to-dos due today...")
}

func (app *Application) createTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	loc, err := time.LoadLocation("Local")
	if err != nil {
		app.serverError(w, err)
		return
	}

	id, err := app.todos.Insert(3, time.Date(2022, 7, 31, 0, 0, 0, 0, loc), "Move into Cinnamon College!", "Remember to pack things and move in before 5 p.m.")
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/todo?id=%d", id),
		http.StatusSeeOther)
}
