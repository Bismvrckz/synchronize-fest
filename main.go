package main

import (
	"doku-api/controllers"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {

	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/", controllers.CreateTicket)
	http.HandleFunc("/create-ticket", controllers.CreateTicket)

	FileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", FileServer))
	log.Print("Server berjalan di port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
