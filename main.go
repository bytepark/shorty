package main

import (
    "fmt"
	"net/http"
    "gopkg.in/flosch/pongo2.v3"
    "github.com/shaoshing/train"
)

type Post struct {
	Url     string
	Comment string
}

var templates = map[string]string {
    "posts":   "html/posts.html",
    "newpost": "html/newpost.html",
    "docs":    "html/docs.html",
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    tpl := pongo2.Must(pongo2.FromFile(templates[tmpl]))
    err := tpl.ExecuteWriter(pongo2.Context{
        "javascript_tag": train.JavascriptTag,
        "stylesheet_tag": train.StylesheetTag,
    }, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerListPosts(writer http.ResponseWriter, request *http.Request) {
	posts := []Post{{Url: "foo", Comment: "Bar"}}
	renderTemplate(writer, "posts", posts)
}

func handlerNewPost(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "newpost", nil)
}

func handlerDocs(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "docs", nil)
}

func handlerShortLink(writer http.ResponseWriter, request *http.Request) {
}

func main() {
    train.ConfigureHttpHandler(nil)

	http.HandleFunc("/posts", handlerListPosts)
	http.HandleFunc("/newpost", handlerNewPost)
	http.HandleFunc("/docs", handlerDocs)
	http.HandleFunc("/(pattern)", handlerShortLink)
	http.HandleFunc("/", handlerListPosts)

    fmt.Println("Listening to localhost:8080")
	err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}
