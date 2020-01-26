package helpers

import (
	"fmt"
	"gopkg.in/go-playground/validator.v8"
)

//ParseError for parse the data to better export in array
func ParseError(err error) []string {
	if err == nil {
		return nil
	}

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}

	var errors []string
	for _, e := range ve {
		if e.Tag == "required" {
			errors = append(errors, fmt.Sprintf("%s is required", e.Field))
		} else {
			errors = append(errors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", e.Field, e.Tag))
		}
	}

	return errors
}
