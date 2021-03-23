package commands

import (
	labels "TDBackend/localization"
	"TDBackend/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Get func for GET RESTful API
func Get(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.Header().Set(labels.HTTPContentTypeKey, labels.HTTPContentTypeValue)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
	logger.AppLogger(labels.Debug, "Result of GET func", time.Since(t).Nanoseconds(), labels.Code+"|"+fmt.Sprintf("%v", http.StatusOK))
}

// Post func for POST RESTful API
func Post(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.Header().Set(labels.HTTPContentTypeKey, labels.HTTPContentTypeValue)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
	logger.AppLogger(labels.Debug, "Result of POST func", time.Since(t).Nanoseconds(), labels.Code+"|"+fmt.Sprintf("%v", http.StatusCreated))
}

// Put func for PUT RESTful API
func Put(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.Header().Set(labels.HTTPContentTypeKey, labels.HTTPContentTypeValue)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
	logger.AppLogger(labels.Debug, "Result of PUT func", time.Since(t).Nanoseconds(), labels.Code+"|"+fmt.Sprintf("%v", http.StatusAccepted))
}

// Delete func for DELETE RESTful API
func Delete(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.Header().Set(labels.HTTPContentTypeKey, labels.HTTPContentTypeValue)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
	logger.AppLogger(labels.Debug, "Result of DELETE func", time.Since(t).Nanoseconds(), labels.Code+"|"+fmt.Sprintf("%v", http.StatusOK))

}

// RunServer func for running HTTP server
func RunServer() error {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", Get).Methods(http.MethodGet)
	api.HandleFunc("", Post).Methods(http.MethodPost)
	api.HandleFunc("", Put).Methods(http.MethodPut)
	api.HandleFunc("", Delete).Methods(http.MethodDelete)

	return http.ListenAndServe(":8080", r)
}
