package handlers

import (
	"go_microservice/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	//fetches the products from the datastore
	lp := data.GetProducts()

	//as we are sending products to user we need to covert lp(go) into json to write in response
	/* d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marhsal json", http.StatusInternalServerError)
	}  */
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

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "cant decode from json", http.StatusBadRequest)
	}

	//adding to datastore
	data.AddProduct(prod)

	//just printing in servers log for confirmation
	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) updateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	//getting update details from req body
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "cant decode from json", http.StatusBadRequest)
	}

	//expect the id in uri
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) != 1 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	if len(g[0]) != 2 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(rw, "Invalid Index", http.StatusBadRequest)
	}

	//updating in datastore
	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Error not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error not found", http.StatusInternalServerError)
		return
	}
}
