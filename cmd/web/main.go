package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog: infoLog,
		errLog:  errorLog,
	}

	err := http.ListenAndServe(":"+os.Getenv("PORT"), app.routes())
	log.Fatal(err)
}
