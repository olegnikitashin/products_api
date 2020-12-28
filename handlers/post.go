package handlers

import (
	"net/http"

	"github.com/olegnikitashin/products_api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//  200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.log.Printf("[POST] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
