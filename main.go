package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var userData UserData

type UserData struct {
	FirstName   string
	LastName    string
	DateOfBirth string
	Sex         string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/templates/init", initHandler)
	http.HandleFunc("/templates/treatment", treatmentHandler)
	http.HandleFunc("/templates/display", displayHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func initHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "init", nil)
}

func treatmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userData = UserData{
			FirstName:   r.FormValue("first_name"),
			LastName:    r.FormValue("last_name"),
			DateOfBirth: r.FormValue("date_of_birth"),
			Sex:         r.FormValue("sex"),
		}

		http.Redirect(w, r, "/templates/display", http.StatusSeeOther)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func displayHandler(w http.ResponseWriter, r *http.Request) {
	// Display the user data

	renderTemplate(w, "display", userData)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplFile := fmt.Sprintf("templates/%s.html", tmpl)
	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
