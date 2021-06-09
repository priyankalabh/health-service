package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/api/health", healthHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error occured in listen n serve %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	x := healthResponse{
		Status: "alive",
	}
	data, err := json.Marshal(&x)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
