package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Product struct{
	ID int
	Name string
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware)
	//query
	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		println("endpoint")
	})

	// params
	router.Get("/{productName}", func(res http.ResponseWriter, req *http.Request) {
		param := chi.URLParam(req, "productName")
		res.Write([]byte(param))
	})

	router.Get("/json", func(res http.ResponseWriter, req *http.Request) {
		obj := map[string]string{"message: ": "sucess"}
		render.JSON(res, req, obj)
	})

	router.Post("/product", func(res http.ResponseWriter, req *http.Request){
		var product Product 

		render.DecodeJSON(req.Body, &product)
		product.ID = 10
		render.JSON(res, req, product)
	})



	http.ListenAndServe(":3000", router)
}


func middleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		println("antes")
		next.ServeHTTP(res, req)
		println("depois")
	})
}