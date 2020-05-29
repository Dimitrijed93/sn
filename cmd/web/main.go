package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Dimitried93/sn/pkg/models"
)

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
	snippet models.SnippetModel
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog: infoLog,
		errLog:  errorLog,
	}

	err := http.ListenAndServe(local, app.routes())
	log.Fatal(err)
}
