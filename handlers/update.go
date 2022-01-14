package handlers

import (
	"net/http"
	"strconv"

	"github.com/LucasSaraiva019/coffe-shop-api/data"
	"github.com/gorilla/mux"
)

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Enable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Note found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product Note found", http.StatusInternalServerError)
		return
	}
}
