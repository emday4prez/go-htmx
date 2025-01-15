package main

import (
	"html/template"
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

}
