package config

type Config struct {
	DomainName string
}

func NewConfig(domainName string) Config {
	return Config{
		DomainName: domainName,
	}
}
