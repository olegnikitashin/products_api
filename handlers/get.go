package handlers

import (
	"net/http"

	"github.com/olegnikitashin/products_api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, req *http.Request) {
	p.log.Println("[GET] /products")

	listOfProducts := data.GetProducts()

	err := data.ToJSON(listOfProducts, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listProduct
// Returns a selected product
// responses:
//  200: productResponse
//  404: errorResponse

// ListAll handles GET requests and returns a selected product
func (p *Products) ListSingle(rw http.ResponseWriter, req *http.Request) {
	id := getProductID(req)

	p.log.Println("[GET] /products/{%d}", id)

	product, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrorProductNotFound:
		p.log.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.log.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(product, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
