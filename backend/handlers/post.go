package handlers

import (
	"net/http"

	"github.com/LucasSaraiva019/coffe-shop-api/backend/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	rw.Header().Add("Content-Type", "application/json")
	p.l.Debug("Inserting product: %#v\n", prod)
	p.productDB.AddProduct(prod)
}
