package main

type Config struct {
	domainName string
}

func NewConfig(domainName string) Config {
	return Config{
		domainName: domainName,
	}
}