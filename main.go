package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	templates     *template.Template
)

type ARTISTS struct {
	name string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var err error

	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Serveur démarré sur http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erreur lors du démarrage du serveur :", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func containsRune(slice []rune, r rune) bool {
	for _, v := range slice {
		if v == r {
			return true
		}
	}
	return false
}

func containsString(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}