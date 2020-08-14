package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string "json:\"Title\""
	Desc    string "json:\"desc\""
	Content string "json:\"content\""
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "test title", Desc: "test desc", Content: "test content"},
	}
	fmt.Println("All Articles end point hit!!!")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit!!!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("service started")
	handleRequests()
}
