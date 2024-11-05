package service

import (
	"math/rand/v2"
	"strconv"

	"github.com/twhl75/url-shortener/internal/models"
)

func urlGen(domainName string, urls models.URLs) models.URLs {
	result := models.URLs{}

	code := rand.IntN(1000)

	result.Original = urls.Original
	result.Shortened = domainName + "/" + strconv.Itoa(code)

	return result
}

func idGen(m map[int]models.URLs) int {
	id := 1
	exists := true

	for exists {
		id = rand.IntN(1000)
		_, exists = m[id]
	}

	return id
}
