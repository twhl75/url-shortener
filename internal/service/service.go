package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/twhl75/url-shortener/internal/config"
	"github.com/twhl75/url-shortener/internal/helper"
	"github.com/twhl75/url-shortener/internal/models"
)


type Service struct {
	db  map[int]models.URLs
	cfg config.Config
	log *log.Logger
}

func NewService() Service {
	return Service{
		db:  make(map[int]models.URLs),
		cfg: config.NewConfig("terence.liu"),
		log: log.New(os.Stdout, "LOG:", log.Lshortfile|log.Ldate),
	}
}

func (s *Service) HandleRoot(w http.ResponseWriter, r *http.Request) {
	s.log.Printf("Welcome to url-shortener\n")
}


func (s *Service) CreateURL(w http.ResponseWriter, r *http.Request) {
	urls := models.URLs{}

	err := json.NewDecoder(r.Body).Decode(&urls)
	if err != nil {
		s.log.Printf("Error decoding json: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = helper.Validate(urls)
	if err != nil {
		s.log.Println("Error validating url:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	urlsShortened := urlGen(s.cfg.DomainName, urls)
	id := idGen(s.db)

	// Store in DB
	s.db[id] = urlsShortened

	message := "Get shortened url at localhost:8080/urls/" + strconv.Itoa(id)
	_, err = w.Write([]byte(message))
	if err != nil {
		s.log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *Service) GetShortenedURL(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		s.log.Printf("Error converting id to int: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	val, exists := s.db[id]
	if !exists {
		s.log.Printf("Error getting URLs: url does not exist")
		http.Error(w, "url does not exist", http.StatusBadRequest)
	}

	_, err = w.Write([]byte(val.Shortened))
	if err != nil {
		s.log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *Service) GetAllURLs(w http.ResponseWriter, r *http.Request) {
	jsonStr, err := json.Marshal(s.db)
	if err != nil {
		s.log.Printf("Error marshaling json: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(jsonStr)
	if err != nil {
		s.log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
