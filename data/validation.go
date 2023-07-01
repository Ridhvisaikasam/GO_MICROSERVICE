package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// ValidationError wraps the validators FieldError so we do not
// expose this to out code
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

type Validation struct {
	validate *validator.Validate
}

// creating validator and registering custom validations
func NewValidation() *Validation {
	validate := validator.New()
	//creating custom validation functions and registering with a name
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

// fl indicates the value to be checked is the field value where we use this custom validater
func validateSKU(fl validator.FieldLevel) bool {
	//sku is of format abc-abcd-abcde
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)
	//exactly one sku
	if len(matches) != 1 {
		return false
	}

	return true
}

// the function which actually validates and implements the logic
func (v *Validation) Validate(i interface{}) ValidationErrors {
	//checks if any validation errors are present
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	//converting each validator.fieldError into ValidationError struct
	var returnErrs ValidationErrors
	for _, err := range errs {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}
