package posts

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "log"
    "os"
)

type Post struct {
    Id        int64
    Created   int64
    Updated   int64
    Shortlink string
    Url       string
    Comment   string
}

var MyLogger = log.New(os.Stderr,"LOG: ", log.Ldate|log.Ltime|log.Lshortfile)


func ListPosts() ([]Post){
    db := OpenDb()
    posts := retrieveAllPosts(db)
    defer db.Close()
    return posts
}

func NewPost(url string, comment string) (Post){
    db := OpenDb()
    post := addPost(db, url, comment)
    defer db.Close()

    return post
}


func retrieveAllPosts(db *sql.DB) ([]Post){
    posts := make([]Post, 1)

    rows, err := db.Query("select id, url from posts")
    if err != nil {
        MyLogger.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var id int64
        var shortlink string
        var url string
        var comment string
        rows.Scan(&id, &shortlink, &url, &comment)
        fmt.Println(id, shortlink, url, comment)

        mypost := Post {Id: id, Shortlink: shortlink, Url: url, Comment: comment}
        posts = append(posts, mypost)
    }
    rows.Close()

    return posts
}


func addPost(db *sql.DB, url string, comment string) Post{
    tx, err := db.Begin()
    if err != nil {
        MyLogger.Fatal(err)
    }
    stmt, err := tx.Prepare("insert into posts(url, comment) values(?, ?)")
    if err != nil {
        MyLogger.Fatal(err)
    }
    defer stmt.Close()
        _, err = stmt.Exec(url, comment)
        if err != nil {
            MyLogger.Fatal(err)
        }
    tx.Commit()

    mypost := Post {Url: url, Comment: comment}
    return mypost
}


func OpenDb() (*sql.DB) {
    db, err := sql.Open("sqlite3", "/tmp/shorty.db")
    if err != nil {
        MyLogger.Fatal(err)
    }

    success := createDb(db, err)
    MyLogger.Printf("OpenDB status: %s", success)

    return db
}


func createDb(db *sql.DB, err error) string{
    sqlStmt := `
        CREATE TABLE IF NOT EXISTS posts (
            id integer not null primary key,
            created integer,
            updated integer,
            shortlink TEXT,
            url TEXT,
            comment TEXT
        );
    `
    _, err = db.Exec(sqlStmt)

    if err != nil {
        MyLogger.Printf("%q: %s\n", err, sqlStmt)
        return "Create stmt failed"
    }

    return "Database opened"
}
