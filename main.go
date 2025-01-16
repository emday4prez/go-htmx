package main

import (
	"html/template"
	"net/http"
	"sync"
)

// Data store (in-memory)
var (
	mu       sync.Mutex
	messages []string
)

// Template rendering
var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", indexHandler)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	tmpl.Execute(w, messages)
}
