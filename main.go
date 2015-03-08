package main

import (
	"net/http"
)

func handlerListPosts(writer http.ResponseWriter, request *http.Request) {
}

func handlerNewPost(writer http.ResponseWriter, request *http.Request) {
}

func handlerDocs(writer http.ResponseWriter, request *http.Request) {
}

func handlerShortLink(writer http.ResponseWriter, request *http.Request) {
}

func main() {
	http.HandleFunc("/posts", handlerListPosts)
	http.HandleFunc("/newpost", handlerNewPost)
	http.HandleFunc("/docs", handlerDocs)
	http.HandleFunc("/(pattern)", handlerShortLink)
	http.HandleFunc("/", handlerListPosts)
	http.ListenAndServe(":8080", nil)
}
