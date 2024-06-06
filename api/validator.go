package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/jainam1259/simplebank/util"
)

// Custom validator function for currency validation
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupported(currency)
	}
	return false
}
