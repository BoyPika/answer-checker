package main

import (
	"log"
	"net/http"
	"os"
)

type Response struct {
	Score int    `json:"score"`
	Pdf   string `json:"pdf"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	alg2()
	chem()
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
