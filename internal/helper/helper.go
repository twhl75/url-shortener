package helper

import (
	"github.com/pkg/errors"
	"github.com/twhl75/url-shortener/internal/models"
)

func Validate(structObj interface{}) error {
	switch structObj := structObj.(type) {
	case models.URL:
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
