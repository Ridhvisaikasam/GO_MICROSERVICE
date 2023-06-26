package data

import "testing"

// to test whether the p.Validate() function is working properly or not
func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "faluda",
		Price: 1.00,
		SKU:   "abs-abs-ab",
	}

	//main testing part
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
