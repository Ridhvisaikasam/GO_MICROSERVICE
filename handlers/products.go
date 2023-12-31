// Package classifaction of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"fmt"
	"go_microservice/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// the above thing was reflected by swagger and make file
// includes all server side logics like middleware and handling requests

type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts returns a new products handler with the given logger , validator
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

//handled by ux router
/*func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//handle retrieval //curl default get
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	//handle an addition //curl with data default post
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	//
	if r.Method == http.MethodPut {
		p.updateProducts(rw, r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}*/

//sent to separate file GET.GO
/*func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	//fetches the products from the datastore
	lp := data.GetProducts()

	//as we are sending products to user we need to covert lp(go) into json to write in response
	// d, err := json.Marshal(lp)
	//if err != nil {
	//	http.Error(rw, "Unable to marhsal json", http.StatusInternalServerError)
	//}
	//encoder -- no need to buffer anything to memory ,, no need to allocate memory to data object
	//encoder is marginally faster , difference is seen when we have multiple threads or large json documents
	//Shifted to data.products as a func to keep the handler clean and not create concrete objs

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode to json", http.StatusInternalServerError)
	}

	//internally written using rw in encode function
	//rw.Write(d)
}
*/

//sent to separate file ADD.GO
/*func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	///*prod := &data.Product{}

	//err := prod.FromJSON(r.Body)
	//if err != nil {
	//	http.Error(rw, "cant decode from json", http.StatusBadRequest)
	//}
	//now getting prod from context from prev middleware
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	//adding to datastore
	data.AddProduct(&prod)

	//just printing in servers log for confirmation
	p.l.Printf("Prod: %#v", prod)
}*/

//sent to separate file UPDATE.GO
/*func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	//getting id from uri but with help of gorilla
	vars := mux.Vars(r)
	idString := vars["id"]

	//getting update details from req body
	//prod := &data.Product{}

	//err := prod.FromJSON(r.Body)
	//if err != nil {
	//	http.Error(rw, "cant decode from json", http.StatusBadRequest)
	//}
	//now getting prod from context from prev middleware
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	//expect the id in uri
	//reg := regexp.MustCompile(`/([0-9]+)`)
	//g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	//if len(g) != 1 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	//if len(g[0]) != 2 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	//idString := g[0][1]
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
}*/

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
