package handlers

import (
	"html/template"
	"net/http"
)

type Hrefs struct {
	Link string
	Text string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	hrefs := []Hrefs{
		{Link: "/plus/", Text: "Операция с плюсом"},
		{Link: "/multiplication/", Text: "Операция с умножением"},
		{Link: "/minus/", Text: "Операция с минусом"},
		{Link: "/divide/", Text: "Операция с делением"},
		{Link: "/cos/", Text: "Операция с косинусом"},
		{Link: "/tan/", Text: "Операция с тангенсом"},
	}

	tmpl, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, hrefs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
