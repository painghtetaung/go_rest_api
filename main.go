package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Article{
		Title:   "Test Title",
		Desc:    "Test Description",
		Content: "Test Content",
	}
	// fmt.Fprintf(w, "Endpoint hit: All articles")
	fmt.Println("Endpoint hit: All articles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request for %s", r.Method, r.URL.Path)
	fmt.Fprintf(w, "Home Page endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
