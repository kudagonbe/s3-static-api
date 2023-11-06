package main

import (
	"fmt"
	"net/http"

	"github.com/kudagonbe/s3-static-api/internal/config"
	"github.com/kudagonbe/s3-static-api/internal/storage"
)

func main() {
	cfg := config.Get()
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := storage.PutObject("sample.png"); err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("Hello World!!!"))
}
