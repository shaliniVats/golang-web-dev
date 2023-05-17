package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}
func main() {
	sages := []string{"Gandhi", "Milk", "Budhha", "Jesus"}
	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}
}
