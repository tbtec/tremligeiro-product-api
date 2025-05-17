package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/tbtec/tremligeiro/internal/types/xerrors"
)

var vld = newValidator()

func newValidator() *validator.Validate {
	vld := validator.New()
	return vld
}

func Validate(input any) error {
	err := vld.Struct(input)
	if err == nil {
		return nil
	}
	return adapt(err)
}

// converts validator erros a xerrors to be treated by the application
func adapt(err error) xerrors.ValidationError {
	vErr := xerrors.NewValidationError("Invalid Body")
	for _, valErr := range err.(validator.ValidationErrors) {
		field := extract(valErr.Namespace())
		switch valErr.Tag() {
		case "required", "required_for", "required_with", "required_with_all":
			vErr = vErr.AddField(field, xerrors.ReasonRequiredAttributeMissing)
		default:
			vErr = vErr.AddField(field, xerrors.ReasonTypeInvalidValue)
		}
	}
	return vErr
}

// get the field from the namespace
func extract(namespace string) string {
	split := strings.Split(namespace, ".")
	if len(split) > 1 {
		return strings.Join(split[1:], ".")
	}
	return split[0]
}
