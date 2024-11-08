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

func (c *Controller) Run(){
	router := http.NewServeMux()
	router.HandleFunc("/", c.svc.HandleRoot)
	router.HandleFunc("POST /url",c.svc.CreateURL)
	router.HandleFunc("GET /url", c.svc.GetAllURLs)
	router.HandleFunc("GET /url/{id}", c.svc.GetShortenedURL)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("server failed to listen: %v", err)
	}
}