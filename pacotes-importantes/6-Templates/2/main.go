package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

func main() {
	curso := Curso{"GO", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))

	err := t.Execute(os.Stdout, &curso)
	if err != nil {
		panic(err)
	}

}
