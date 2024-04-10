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
	log.Fatal(http.ListenAndServe(":8088", mux))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	curso := Curso{"GO", 40}
	curso1 := Curso{"Java", 60}
	curso2 := Curso{"Python", 55}
	curso3 := Curso{"C#", 30}
	cursos := Cursos{curso, curso1, curso2, curso3}

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(w, &cursos)
	if err != nil {
		panic(err)
	}
}
