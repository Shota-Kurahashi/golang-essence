package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name string
}

func main() {
	user := User{Name: "Tom"}

	tmpl := `{{.Name}}`
	t := template.Must(template.New("tmpl").Parse(tmpl))

	err := t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}
}
