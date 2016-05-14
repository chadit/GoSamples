package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

// to run from the command line type go run main.go
// Open a browser and navigate to localhost:5555
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content Type", "text/html")
		//	tmpl, err := template.New("test").Parse(doc)
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)

		context := Context{
			[3]string{"Lemon", "Orange", "Apple"},
			"the title",
		}
		templates.Lookup("test").Execute(w, context)

	})
	http.ListenAndServe(":5555", nil)

}

const doc = `
{{template "header" .Title}}
<body>
<h1>List of Fruit</h1>
<ul>
{{range .Fruit}}
<li>{{.}}</li>
{{end}}
</ul>
</body>
{{template "footer"}}
`

const header = `
<!DOCTYPE html>
<html>
<head><title>{{.}}</title></head>
`

const footer = `
</html>
`

// Context - sample
type Context struct {
	Fruit [3]string
	Title string
}
