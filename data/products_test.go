package data

import (
	"testing"
)

// to test whether the p.Validate() function is working properly or not
func TestChecksValidation(t *testing.T) {
	v := NewValidation()
	prod := &Product{
		ID:          0,
		Name:        "faluda",
		Description: "",
		Price:       1.00,
		SKU:         "abs-abs-ab",
	}

	//main testing part
	errs := v.Validate(prod)
	if len(errs) != 0 {
		t.Fatal(errs)
	}
}
