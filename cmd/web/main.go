package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	log.Fatal(err)
}
