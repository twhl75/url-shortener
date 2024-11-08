package models

type URL struct {
	ID        int64  `json:"id"`
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
}
