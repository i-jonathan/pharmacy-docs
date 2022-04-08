package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	rapidoc := http.FileServer(http.Dir("./docs"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", rapidoc))

	port := os.Getenv("PORT")

	_ = http.ListenAndServe(":"+port, router)
}
