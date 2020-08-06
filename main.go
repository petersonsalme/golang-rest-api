package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title   string `json:"titulo"`
	Desc    string `json:"descricao"`
	Content string `json:"conteudo"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "How to build a RESTful api with Go", Desc: "Descrição artigo de Go", Content: "Conteudo"},
	}

	fmt.Fprintln(w, "All Articles:")
	json.NewEncoder(w).Encode(articles)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/articles", allArticles)
	http.ListenAndServe(":8081", nil)
}
