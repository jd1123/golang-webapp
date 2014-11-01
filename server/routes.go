package server

import "net/http"

type HF func(w http.ResponseWriter, r *http.Request)

var routes = map[string]HF{
	"/":             IndexHandler,
	"/dir1/{name}":  DirHandler,
	"/files/":       ShowFilesHandler,
	"/files/{path}": ShowFilesHandler,
}
