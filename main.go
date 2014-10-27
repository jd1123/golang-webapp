package main

import (
	"fmt"
	"net/http"

	"github.com/jd1123/golang-webapp/server"
)

func main() {
	fmt.Println("Listening on port 8081")
	server.RegisterHandlers()
	http.ListenAndServe(":8081", nil)
}
