package main

import (
    "database/sql"
    "fmt"
    "net/http"
    _ "github.com/lib/pq" // PostgreSQL driver
    "html/template"
)

var db *sql.DB

func main() {
    var err error
    // Connect to PostgreSQL database
    connStr := "user=username dbname=newstar sslmode=disable" // update with your credentials
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    http.HandleFunc("/", showForm)
    http.HandleFunc("/register", registerStudent)

    http.ListenAndServe(":8080", nil)
}

func showForm(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, nil)
}

func registerStudent(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        email := r.FormValue("email")
        course := r.FormValue("course")

        _, err := db.Exec("INSERT INTO students(name, email, course) VALUES($1, $2, $3)", name, email, course)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Registration successful for %s!", name)
    }
}
