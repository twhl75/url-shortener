package main

import (
	"math/rand/v2"
	"strconv"
)

func urlGen(u *UrlShortener, urls URLs) URLs {
	result := URLs{}

	code := rand.IntN(1000)

	result.Original = urls.Original
	result.Shortened = u.cfg.domainName + "/" + strconv.Itoa(code)

	return result
}

func idGen(u *UrlShortener) int {
	id := 1
	exists := true

	for exists {
		id = rand.IntN(1000)
		_, exists = u.db[id]
	}

	return id
}
