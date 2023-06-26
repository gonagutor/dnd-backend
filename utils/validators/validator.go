package validators

import (
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func SetupValidator() {
	Validator.RegisterValidation("password", Password)
}
