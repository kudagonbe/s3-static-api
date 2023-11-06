package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kudagonbe/s3-static-api/internal/config"
	"github.com/kudagonbe/s3-static-api/internal/storage"
)

func main() {
	cfg := config.Get()
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil)
}

type putRequest struct {
	Key string `json:"key"`
}

type response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPut:
		putHandler(w, r)
	default:
		errorResponse(w, fmt.Errorf("method %s is not allowed", method))
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	qs := r.URL.Query()
	if len(qs["key"]) == 0 {
		errorResponse(w, fmt.Errorf("key is required"))
		return
	}
	key := qs["key"][0]
	_, err := storage.GetObject(key)
	if err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response{Message: "success"}); err != nil {
		errorResponse(w, err)
		return
	}
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	req := putRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err := storage.PutObject(req.Key); err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response{Message: "success"}); err != nil {
		errorResponse(w, err)
		return
	}
}

func errorResponse(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	res, err := json.Marshal(response{Message: e.Error()})
	if err != nil {
		s := `{"message": "error"}`
		if _, err := w.Write([]byte(s)); err != nil {
			log.Println(err)
		}
		return
	}
	if _, err := w.Write([]byte(res)); err != nil {
		log.Println(err)
	}
}
