package main

import (
	"math/rand/v2"
	"strconv"
)

func urlGen(urls *URLs, u *UrlShortener) {
	code := rand.IntN(1000)
	urls.Shortened = u.cfg.domainName + "/" + strconv.Itoa(code)
}

func idGen(urls *URLs, u *UrlShortener) int {
	id := 1
    exists := true

    for exists {
        id = rand.IntN(1000)
        _, exists = u.db[id]
    }
    u.db[id] = *urls

	return id
}