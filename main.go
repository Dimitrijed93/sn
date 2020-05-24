package main

import (
	"fmt"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HEROKUUUUUUUUUUUUUUUU"))
}

func main() {

	fmt.Println("DSDD")
	http.Handle("/handle", http.HandlerFunc(handlerFunc))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
