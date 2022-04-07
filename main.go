package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//redoc := http.FileServer(http.Dir("./redoc"))
	rapidoc := http.FileServer(http.Dir("./docs"))
	//yml := http.FileServer(http.Dir("./"))

	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", rapidoc))
	//router.PathPrefix("/docs/redoc").Handler(http.StripPrefix("/docs/redoc", redoc))

	_ = http.ListenAndServe(":7000", router)
}
