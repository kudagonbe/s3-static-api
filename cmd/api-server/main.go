package main

import (
	"fmt"
	"net/http"

	"github.com/kudagonbe/s3-static-api/internal/config"
)

func main() {
	cfg := config.Get()
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!"))
}
