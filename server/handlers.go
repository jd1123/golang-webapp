package server

import (
	"io/ioutil"
	"net/http"
	"path"
)

var homedir string = GetCWD()
var indexPage string = "index.html"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fn := path.Join(homedir, "files", indexPage)
	body, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

func DirHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my dir handler"))
}
