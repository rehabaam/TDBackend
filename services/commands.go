package commands

import (
	"net/http"

	"github.com/gorilla/mux"
)

// getPartners func for getting TriDubai annually partners
func getPartners(w http.ResponseWriter, r *http.Request) {
	_, _ = getData("Partners", w, r)
}

// getDeals func for getting TriDubai annually deals
func getDeals(w http.ResponseWriter, r *http.Request) {
	_, _ = getData("Deals", w, r)
}

// getSessions func for getting TriDubai weekly seesions
func getSessions(w http.ResponseWriter, r *http.Request) {
	_, _ = getData("Sessions", w, r)
}

// getKit func for getting TriDubai Triathlon kit
func getKit(w http.ResponseWriter, r *http.Request) {
	_, _ = getData("Kit", w, r)
}

// getKit func for getting TriDubai Triathlon kit
func getFAQs(w http.ResponseWriter, r *http.Request) {
	_, _ = getData("FAQs", w, r)
}

// serveImage func for serving TriDubai Session images
func serveImage(w http.ResponseWriter, r *http.Request) {
	_, _ = getImage(w, r)
}

// RunServer func for running HTTP server
func RunServer() error {

	loadFileToMemory()
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/partners", getPartners).Methods(http.MethodGet)
	api.HandleFunc("/deals", getDeals).Methods(http.MethodGet)
	api.HandleFunc("/sessions", getSessions).Methods(http.MethodGet)
	api.HandleFunc("/kit", getKit).Methods(http.MethodGet)
	api.HandleFunc("/faqs", getFAQs).Methods(http.MethodGet)
	api.HandleFunc("/img/{name}", serveImage).Methods(http.MethodGet)

	return http.ListenAndServe(":8080", r)
}
