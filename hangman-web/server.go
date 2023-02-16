package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	Hangman ".git/hangman-web/templates"
)

func main() {

	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileserver))

	http.HandleFunc("/", Handler)
	http.HandleFunc("/credits.html", Handler)
	http.HandleFunc("/about.html", Handler)
	http.HandleFunc("/login.html", Handler)
	http.HandleFunc("/jeu.html", Handler)
	http.HandleFunc("/setup", Handler)
	http.HandleFunc("/waiting", Handler)
	http.HandleFunc("/win.html", Handler)
	http.HandleFunc("/lose.html", Handler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

var Data Hangman.HangManData

func Handler(w http.ResponseWriter, r *http.Request) {

	var templateshtml = template.Must(template.ParseGlob("./static/*.html"))
	switch r.URL.Path {
	case "/":
		// accueil
		templateshtml.ExecuteTemplate(w, "index.html", "")
	case "/about.html":
		// about page
		templateshtml.ExecuteTemplate(w, "about.html", "")
	case "/credits.html":
		// credit page
		templateshtml.ExecuteTemplate(w, "credits.html", "")
	case "/login.html":
		// login page
		templateshtml.ExecuteTemplate(w, "login.html", Data)
	case "/setup":
		// setup page
		Data = Hangman.Setup()
		Data.Name = r.FormValue("Name")
		http.Redirect(w, r, "jeu.html", http.StatusFound)
	case "/jeu.html":
		// game page
		if Hangman.Win(Data) {
			http.Redirect(w, r, "win.html", http.StatusFound)
		} else if Hangman.Lose(Data) {
			http.Redirect(w, r, "lose.html", http.StatusFound)
		} else {
			templateshtml.ExecuteTemplate(w, "jeu.html", Data)
		}
	case "/waiting":
		// refreshing page
		Data = Hangman.Jeu(Data, r.FormValue("lettre"))
		http.Redirect(w, r, "jeu.html", http.StatusFound)
	case "/win.html":
		// winning page
		templateshtml.ExecuteTemplate(w, "win.html", Data)
	case "/lose.html":
		// losing page
		templateshtml.ExecuteTemplate(w, "lose.html", Data)
	}
}
