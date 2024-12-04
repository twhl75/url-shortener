package service

import (
	"math/rand/v2"
	"strconv"

	"github.com/twhl75/url-shortener/internal/models"
)

func shortenedGen(urls models.URL) models.URL {
	result := models.URL{}

	code := rand.IntN(1000)

	result.Original = urls.Original
	result.Shortened = strconv.Itoa(code)

	return result
}
