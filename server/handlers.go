package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"text/template"

	"github.com/gorilla/mux"
)

var homedir string = GetCWD()
var indexPage string = "index.html"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile(TemplatePath("index.html"))
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

func DirHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	type Person struct {
		Name string
	}
	p := Person{Name: vars["name"]}
	tmpl, err := template.New("dir.html").ParseFiles(TemplatePath("dir.html"))
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

func ShowFilesHandler(w http.ResponseWriter, r *http.Request) {
	//Check for mux variables
	vars := mux.Vars(r)

	// structs for template
	type S struct {
		Files []string
	}
	baseDir := "/home/jw/"
	useDir := baseDir
	if vars["path"] != "" {
		useDir = path.Join(baseDir, vars["path"])
	}

	contents, err := ioutil.ReadDir(useDir)
	if err != nil {
		panic(err)
	}

	fmt.Println(contents[0].Name())
	OutputStruct := S{Files: PackFiles(contents)}

	tmpl, err := template.New("files.html").ParseFiles(TemplatePath("files.html"))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, OutputStruct)
	if err != nil {
		panic(err)
	}
}
