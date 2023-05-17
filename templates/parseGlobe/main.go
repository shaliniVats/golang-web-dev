package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("tmpls/*"))
}

func main() {
	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	readtmplFile("one.gohtml", err)
	readtmplFile("two.gohtml", err)
	readtmplFile("vespa.gohtml", err)
}

func readtmplFile(fname string, err error) {
	//fmt.Println(fname)
	err = tmpl.ExecuteTemplate(os.Stdout, fname, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
