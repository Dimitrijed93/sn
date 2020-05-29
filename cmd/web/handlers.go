package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/Dimitried93/sn/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
	}
	w.Write([]byte("Hello from Snippetbox"))
}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func (app *application) getAllSnippets(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get(getAll)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		app.serverError(w, err)
	}
	log.Println(string(body))
}

func (app *application) createSnippetSave(w http.ResponseWriter, r *http.Request) {

	var snippet = &models.Snippet{
		ID:      5,
		Title:   "From web app",
		Created: time.Now(),
	}

	json, _ := json.Marshal(snippet)

	res, err := http.Post(insert, "applicaton/json", bytes.NewBuffer(json))

	if err != nil {
		app.serverError(w, err)
	}

	fmt.Println(res.Body)

	// http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
