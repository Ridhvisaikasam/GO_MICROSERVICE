package handlers

import (
	"go_microservice/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Return a list of products
// responses:
//   200: productsResponse

// Get Products returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
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

	err := data.ToJSON(lp, rw)
	if err != nil {
		http.Error(rw, "Unable to encode to json", http.StatusInternalServerError)
	}

	//internally written using rw in encode function
	//rw.Write(d)
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) GetProductByID(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrorProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
