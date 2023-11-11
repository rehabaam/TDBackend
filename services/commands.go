package commands

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var (
	server *http.Server
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
func StartServer() error {

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/partners", getPartners).Methods(http.MethodGet)
	api.HandleFunc("/deals", getDeals).Methods(http.MethodGet)
	api.HandleFunc("/sessions", getSessions).Methods(http.MethodGet)
	api.HandleFunc("/kit", getKit).Methods(http.MethodGet)
	api.HandleFunc("/faqs", getFAQs).Methods(http.MethodGet)
	api.HandleFunc("/img/{name}", serveImage).Methods(http.MethodGet)

	loadFileToMemory()

	server = &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
		Handler:           r,
	}

	ctx := context.Background()
	StopServer(ctx, server)

	return server.ListenAndServe()
}

// Graceful shutdown for HTTP server if Interrupt signal captured from OS
func StopServer(ctx context.Context, srv *http.Server) {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			break
		case <-ctx.Done():
			break
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()
}
