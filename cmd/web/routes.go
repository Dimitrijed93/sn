package main

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet", http.HandlerFunc(app.showSnippet))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippetSave))
	mux.Get("/snippet/get", http.HandlerFunc(app.getAllSnippets))

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, customHeader)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}
