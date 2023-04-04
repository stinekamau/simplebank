package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/stinekamau/simplebank/utils"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {

	if currency, ok := fl.Field().Interface().(string); ok {
		// check currency is supported
		return utils.IsSupportedCurrency(currency)
	}

	return false 

}
