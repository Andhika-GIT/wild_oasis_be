package apperror

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ExtractValidationError(err error) []string {
	var messages []string

	errs, ok := err.(validator.ValidationErrors)

	if ok {
		for _, e := range errs {
			fieldName := e.Field()
			tag := e.Tag()
			messages = append(messages, fmt.Sprintf("Field %s is %s", fieldName, tag))
		}
	}

	return messages
}
