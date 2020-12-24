package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olegnikitashin/products_api/data"
)

type Products struct {
	log *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request) {
	p.log.Println("[GET] /products")

	listOfProducts := data.GetProducts()

	err := listOfProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter, req *http.Request) {
	p.log.Println("[POST] /products")

	product := req.Context().Value(KeyProduct{}).(data.Product)

	p.log.Printf("Product: %#v", product)
	data.AddProduct(&product)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	p.log.Println("[PUT] /products")
	product := req.Context().Value(KeyProduct{}).(data.Product)

	// product := &data.Product{}

	// err = product.FromJSON(req.Body)
	// if err != nil {
	// 	http.Error(rw, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	err = data.UpdateProduct(id, &product)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		product := data.Product{}

		err := product.FromJSON(req.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), KeyProduct{}, product)
		request := req.WithContext(ctx)

		next.ServeHTTP(rw, request)
	})
}
