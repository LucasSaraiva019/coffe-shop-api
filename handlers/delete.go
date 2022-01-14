package handlers

import (
	"net/http"
	"strconv"

	"github.com/LucasSaraiva019/coffe-shop-api/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns no Content
// responses:
// 	200: noContent

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Enable to convert id", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle DELETE Product", id)

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Note found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product Note found", http.StatusInternalServerError)
		return
	}

}
