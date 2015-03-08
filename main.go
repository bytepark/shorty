package main

import (
    "github.com/shaoshing/train"
    "html/template"
    "net/http"
    "fmt"
    posts "github.com/bytepark/shorty/posts"
)

var templates = template.Must(template.ParseFiles("html/posts.tmpl.html", "html/newpost.tmpl.html", "html/docs.tmpl.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handlerListPosts(writer http.ResponseWriter, request *http.Request) {
    myposts := posts.ListPosts()
    renderTemplate(writer, "posts.tmpl.html", myposts)
}

func handlerNewPost(writer http.ResponseWriter, request *http.Request) {
    url := request.FormValue("url")
    fmt.Println("new url", url)

    mypost := posts.NewPost(url, "testlink")
    renderTemplate(writer, "newpost.tmpl.html", mypost)
}

func handlerDocs(writer http.ResponseWriter, request *http.Request) {
    renderTemplate(writer, "docs.tmpl.html", nil)
}

func handlerShortLink(writer http.ResponseWriter, request *http.Request) {
}

func main() {
    http.HandleFunc("/posts", handlerListPosts)
    http.HandleFunc("/newpost", handlerNewPost)
    http.HandleFunc("/docs", handlerDocs)
    http.HandleFunc("/(pattern)", handlerShortLink)
    http.HandleFunc("/", handlerListPosts)
    train.ConfigureHttpHandler(nil)
    http.ListenAndServe(":8080", nil)
}
