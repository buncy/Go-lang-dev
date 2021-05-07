package main

import (
	"html/template"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	todos := ToDo{
		User: "tim",
		List: []entry{
			{
				Name: "todo1",
				Done: false,
			},
			{
				Name: "todo2",
				Done: false,
			},
		},
	}

	// Files are provided as a slice of strings.
	// paths := []string{
	// 	"test.tmpl",
	// }

	t := template.Must(template.New("test.tmpl").ParseFiles("test.tmpl"))
	err := t.Execute(os.Stdout, todos)
	if err != nil {
		panic(err)
	}
}
