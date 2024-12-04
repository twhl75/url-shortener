package service

import (
	"database/sql"
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
	db  *sql.DB
	cfg config.Config
	log *log.Logger
}

func NewService(database *sql.DB) Service {
	return Service{
		db:  database,
		cfg: config.NewConfig("terence.liu"),
		log: log.New(os.Stdout, "LOG:", log.Lshortfile|log.Ldate),
	}
}

func (s *Service) HandleRoot(w http.ResponseWriter, r *http.Request) {
	s.log.Printf("Welcome to url-shortener\n")
}


func (s *Service) CreateURL(w http.ResponseWriter, r *http.Request) {
	url := models.URL{}

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		s.log.Printf("Error decoding json: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = helper.Validate(url)
	if err != nil {
		s.log.Println("Error validating url:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url = shortenedGen(url)

	result, err := s.db.Exec("INSERT INTO url (original, shortened) VALUES (?, ?)", url.Original, url.Shortened)
	if err != nil {
		s.log.Println("Error querying:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		s.log.Println("Error retrieving id:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	

	message := "Get shortened url at localhost:8080/urls/" + strconv.Itoa(int(id))
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

	url := models.URL{}

	row := s.db.QueryRow("SELECT * FROM url WHERE ID = ?", id)
	if err := row.Scan(&url.ID, &url.Original, &url.Shortened); err != nil {
		s.log.Printf("Error scanning queried row: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte(s.cfg.DomainName + "/" + url.Shortened))
	if err != nil {
		s.log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *Service) GetAllURLs(w http.ResponseWriter, r *http.Request) {
	urls := []models.URL{}

	rows, err := s.db.Query("SELECT * FROM url")
	if err != nil {
		s.log.Printf("Error querying: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for rows.Next() {
		url := models.URL{}
		if err := rows.Scan(&url.ID, &url.Original, &url.Shortened); err != nil {
			s.log.Printf("Error scanning queried rows: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		urls = append(urls, url)
	}

	jsonStr, err := json.Marshal(urls)
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
