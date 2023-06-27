package handlers

import (
	"go_microservice/data"
	"net/http"
)

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	/*prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "cant decode from json", http.StatusBadRequest)
	}*/
	//now getting prod from context from prev middleware
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	//adding to datastore
	data.AddProduct(&prod)

	//just printing in servers log for confirmation
	p.l.Printf("Prod: %#v", prod)
}
