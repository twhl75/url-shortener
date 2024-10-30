package main

import (
	"github.com/pkg/errors"
)

func validate(structObj interface{}) error {
	switch structObj := structObj.(type) {
	case URLs:
		return validateEmptyURL(structObj.Original)
	default:
		return errors.New("unsupported type")
	}
}

func validateEmptyURL(url string) error {
	if url == "" {
		return errors.New("url cannot be empty")
	}
	return nil
}
