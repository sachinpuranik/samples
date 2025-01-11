package main

import (
	"fmt"
	"net/http"
	// "os"
	"log"
	"net"
	"path/filepath"
	"strings"

	// "fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
)

func courseHandler(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	splitPath := strings.Split(uri, "/")[1:]
	log.Println("splitPath :", splitPath)

	newPath := ""
	for i, s := range splitPath {

		newPath = newPath + "/" + s
		if i == 1 {
			newPath = newPath + "/build"
		}
	}

	if filepath.Ext(newPath) == "" {
		newPath = newPath + "/index.html"
	}

	newPath = "." + newPath

	log.Println("new path:", newPath)

	http.ServeFile(w, r, newPath)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Home</title>
		</head>
		<body>
			<h1>├ Gorilla! Home あ</h1>
		</body>
	</html>
	`
	w.Write([]byte(html))
}

func main() {
	r := mux.NewRouter()

	r.Use(middleware.Logger)
	r.HandleFunc("/", homeHandler)
	r.PathPrefix("/courses").Handler(http.HandlerFunc(courseHandler))

	// http.ListenAndServe("127.0.0.1:3030", r)
	l, _ := net.Listen("tcp", "0.0.0.0:3030")
	fmt.Printf("server listening at %s", l.Addr().String())

	// http.ListenAndServe("127.0.0.1:3030", r)
	http.Serve(l, r)
}
