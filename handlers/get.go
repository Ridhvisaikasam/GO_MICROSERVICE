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

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode to json", http.StatusInternalServerError)
	}

	//internally written using rw in encode function
	//rw.Write(d)
}
