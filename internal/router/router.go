package router

import (
	"github.com/adamisrael/gedcom-api/internal/endpoints"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/", endpoints.IndexHandler)

	r.HandleFunc("/upload", endpoints.UploadHandler).Methods("POST")
	// r.HandleFunc("/upload/{gedcomid}", handler).Methods("DELETE")

	// TODO: use separate handlers
	// TODO: Make endpoints namespace, so gedcom-api/endpoints/individual.go would have the Individual handler(s)
	// TODO: Add statistics/ handlers, i.e., surnames, names, places, etc.
	// r.HandleFunc("/individual", IndividualHandler).Methods("GET")
	r.HandleFunc("/individual/{gedcomid}", endpoints.IndividualHandler).Methods("GET")

	r.HandleFunc("/individual/{gedcomid}/{id:[0-9]+}", endpoints.IndividualHandler).Methods("GET")

	// Get a list of all surnames in the gedcom and how frequently they appear
	r.HandleFunc("/surname/{gedcomid}", endpoints.SurnameHandler)

	return r
}
