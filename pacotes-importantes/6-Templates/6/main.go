package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

type Cursos []Curso

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8090", mux))

}

func ToUpperString(s string) string {
	return strings.ToUpper(s)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	cursos := Cursos{{"Go", 40}, {"Java", 60}, {"Python", 55}, {"C#", 30}}
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpperString": ToUpperString})
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(w, &cursos)
	if err != nil {
		panic(err)
	}
}
