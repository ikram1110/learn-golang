package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	id := vars["id"]

	var edtArticle Article
	json.Unmarshal(reqBody, &edtArticle)
	json.NewEncoder(w).Encode(edtArticle)

	for index, article := range Articles {
		if article.Id == id {
			fmt.Println("index Hit: ", index)
			Articles[index].Title = edtArticle.Title
			Articles[index].Desc = edtArticle.Desc
			Articles[index].Content = edtArticle.Content
		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{
			Id:      "1",
			Title:   "Hello",
			Desc:    "Article Description",
			Content: "Article Content",
		},
		Article{
			Id:      "2",
			Title:   "Hello 2",
			Desc:    "Article Description 2",
			Content: "Article Content 2",
		},
	}
	handleRequests()
}
