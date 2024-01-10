package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main(){
	r := chi.NewRouter()

	//query
	r.Get("/", func(w http.ResponseWriter, req *http.Request){
		product := req.URL.Query().Get("product")
		id := req.URL.Query().Get("id")
		w.Write([]byte(product + id))
	})

	// params
	r.Get("/{productName/id}", func(w http.ResponseWriter, req *http.Request){
		param := chi.URLParam(req, "productName", "id")
		w.Write([]byte(param))
	})

	http.ListenAndServe(":3000", r)
}