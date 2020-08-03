package main

import (
	"net/http"
	"text/template"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	name := req.URL.Query().Get("name")

	tmpl, _ := template.New("data").Parse("<h1>Hello {{ .}}</h1>")
	tmpl.Execute(res, name)

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
