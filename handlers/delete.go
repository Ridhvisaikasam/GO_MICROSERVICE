package handlers

import (
	"go_microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProducts
// Nothing returns in the response
// responses:
//   201: noContent

// Delete Products deletes a product from the data store
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Products")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
