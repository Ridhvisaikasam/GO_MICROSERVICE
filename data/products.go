// kind of database ,, so as to abstract all working just like a database
package data

import (
	"fmt"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

// so that we can add function
type Products []*Product

// always to try to abstract the logic where data is coming from
func GetProducts() Products {
	return productList
}

func GetProductByID(id int) (Product, error) {
	prod, _, err := findProductByID(id)
	if err != nil {
		return *prod, err
	}

	return *prod, nil
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(p *Product) error {
	_, pos, err := findProductByID(p.ID)
	if err != nil {
		return err
	}

	productList[pos] = p
	return nil
}

func DeleteProduct(id int) error {
	_, pos, err := findProductByID(id)
	if err != nil {
		return err
	}

	productList[pos] = productList[len(productList)-1]
	productList = productList[:len(productList)-1]

	return nil
}

func findProductByID(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrorProductNotFound
}

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrorProductNotFound = fmt.Errorf("Product not found")

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
