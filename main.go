package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type UrlShortener struct {
	db  map[int]URLs
	cfg Config
	log *log.Logger
}

type URLs struct {
	Original  string `json:"original"`
	Shortened string
}

func main() {
	urlService := NewUrlShortener()
	urlService.log.Println("Starting server...")

	router := http.NewServeMux()
	router.HandleFunc("/", urlService.handleRoot)
	router.HandleFunc("POST /urls", urlService.createURL)
	router.HandleFunc("GET /urls", urlService.getAllURLs)
	router.HandleFunc("GET /urls/{id}", urlService.getShortenedURL)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("server failed to listen: %v", err)
	}
}

func NewUrlShortener() UrlShortener {
	return UrlShortener{
		db:  make(map[int]URLs),
		cfg: NewConfig("terence.liu"),
		log: log.New(os.Stdout, "LOG:", log.Lshortfile|log.Ldate),
	}
}

func (u *UrlShortener) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to url-shortener\n")
}

func (u *UrlShortener) createURL(w http.ResponseWriter, r *http.Request) {
	urls := URLs{}

	err := json.NewDecoder(r.Body).Decode(&urls)
	if err != nil {
		u.log.Printf("Error decoding json: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate(urls)
	if err != nil {
		u.log.Println("Error validating url:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	urlsShortened := urlGen(u, urls)
	id := idGen(u)

	// Store in DB
	u.db[id] = urlsShortened

	message := "Get shortened url at localhost:8080/urls/" + strconv.Itoa(id)
	_, err = w.Write([]byte(message))
	if err != nil {
		u.log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (u *UrlShortener) getShortenedURL(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		u.log.Printf("Error converting id to int: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	val, exists := u.db[id]
	if !exists {
		u.log.Printf("Error getting URLs: url does not exist")
		http.Error(w, "url does not exist", http.StatusBadRequest)
	}

	_, err = w.Write([]byte(val.Shortened))
	if err != nil {
		u.log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (u *UrlShortener) getAllURLs(w http.ResponseWriter, r *http.Request) {
	jsonStr, err := json.Marshal(u.db)
	if err != nil {
		u.log.Printf("Error marshaling json: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(jsonStr)
	if err != nil {
		u.log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
