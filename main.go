package main

import (
	"log"
	"net/http"

	"github.com/cojoj/analyzer/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/analyze", handler.Analyze)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
