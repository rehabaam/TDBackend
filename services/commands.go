package commands

import (
	"net/http"

	"github.com/gorilla/mux"
)

// getPartners func for getting TriDubai annually partners
func getPartners(w http.ResponseWriter, r *http.Request) {
	readFileData("Partners", w, r)
}

// getDeals func for getting TriDubai annually deals
func getDeals(w http.ResponseWriter, r *http.Request) {
	readFileData("Deals", w, r)
}

// getSessions func for getting TriDubai weekly seesions
func getSessions(w http.ResponseWriter, r *http.Request) {
	readFileData("Sessions", w, r)
}

// RunServer func for running HTTP server
func RunServer() error {
	r := mux.NewRouter()

	api := r.PathPrefix("/TriDubai/api/v1").Subrouter()
	api.HandleFunc("/partners", getPartners).Methods(http.MethodGet)
	api.HandleFunc("/deals", getDeals).Methods(http.MethodGet)
	api.HandleFunc("/sessions", getSessions).Methods(http.MethodGet)

	return http.ListenAndServe(":8080", r)
}
