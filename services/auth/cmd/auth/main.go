package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() 
	sugar := logger.Sugar()
	sugar.Infof("Failed to fetch URL: %s", "url")

	r := mux.NewRouter()

	port := "8080"
	sugar.Infof("Starting server on port %s", port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		sugar.Fatalf("Server failed: %v", err)
	}

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		sugar.Info("Hello endpoint hit")
		w.Write([]byte("Hello from service!"))
	}).Methods("GET")
}
