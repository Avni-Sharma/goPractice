package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	t, err := template.ParseFiles("gotempl.tmpl")
	if err != nil {
		log.Fatal("Could not parse files")
	}
	err = t.Execute(os.Stdout, "nil\n")
	if err != nil {
		log.Fatal("could not execute templates 	")
	}
	err = t.ExecuteTemplate(os.Stdout, "gotempl.tmpl", "Avni\n")
	if err != nil {
		log.Fatal("could not execute templates")
	}
}
