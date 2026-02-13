package main

import (
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/iliamunaev/go-agg/internal/agg"
)

func main() {
	a := agg.New(&http.Client{}, 8, 2*time.Second)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/aggregate", func(w http.ResponseWriter, r *http.Request) {
		urls := r.URL.Query()["u"]
		if err := a.Aggregate(r.Context(), urls); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	go func() {
		log.Println("pprof on http://localhost:6060/debug/pprof/")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}