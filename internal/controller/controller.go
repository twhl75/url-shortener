package controller

import (
	"log"
	"net/http"

	"github.com/twhl75/url-shortener/internal/service"
)

type Controller struct {
	svc service.Service
}

func New(svc service.Service) Controller {
	return Controller{
		svc: svc,
	}
}

func (c *Controller) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", c.svc.HandleRoot)
	mux.HandleFunc("POST /url", c.svc.CreateURL)
	mux.HandleFunc("GET /url", c.svc.GetAllURLs)
	mux.HandleFunc("GET /url/{id}", c.svc.GetShortenedURL)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("server failed to listen: %v", err)
	}
}
