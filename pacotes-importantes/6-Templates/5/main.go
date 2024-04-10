package main

import (
	"html/template"
	"log"
	"net/http"
)

type Curso struct {
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

type Cursos []Curso

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8089", mux))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	cursos := Cursos{{"GO", 40}, {"Java", 60}, {"Python", 55}, {"C#", 30}}
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))

	err := t.Execute(w, &cursos)
	if err != nil {
		panic(err)
	}
}
