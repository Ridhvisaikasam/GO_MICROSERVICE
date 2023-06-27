package handlers

import (
	"go_microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	//getting id from uri but with help of gorilla
	vars := mux.Vars(r)
	idString := vars["id"]

	//getting update details from req body
	/*prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "cant decode from json", http.StatusBadRequest)
	}*/
	//now getting prod from context from prev middleware
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	//expect the id in uri
	/*reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) != 1 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	if len(g[0]) != 2 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	idString := g[0][1]*/
	//convert string to int
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(rw, "Invalid Index", http.StatusBadRequest)
	}

	//updating in datastore
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Error not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error not found", http.StatusInternalServerError)
		return
	}
}
