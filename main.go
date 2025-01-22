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
	http.HandleFunc("/add", addHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	tmpl.Execute(w, messages)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		msg := r.FormValue("message")
		mu.Lock()
		messages = append(messages, msg)
		mu.Unlock()

		// Respond with the new item for HTMX to swap in
		w.Write([]byte("<li>" + msg + ` <button hx-post="/delete" hx-vals='{"message":"` + msg + `"}'>Delete</button></li>`))
	}
}
