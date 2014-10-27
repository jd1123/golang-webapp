package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type HF func(w http.ResponseWriter, r *http.Request)

var routes = map[string]HF{
	"/":      IndexHandler,
	"/dir1/": DirHandler,
}

func RegisterHandlers() {
	r := mux.NewRouter()
	for k, v := range routes {
		r.HandleFunc(k, v)
	}
	http.Handle("/", r)
}

func GetCWD() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
