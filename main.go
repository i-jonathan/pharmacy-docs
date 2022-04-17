package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	rapidoc := http.FileServer(http.Dir("./docs"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", rapidoc))

	port := os.Getenv("PORT")
	// port := "7000"

	_ = http.ListenAndServe(":"+port, router)
}
