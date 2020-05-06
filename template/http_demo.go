package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/get/", handleGet)
	http.HandleFunc("/post/", handlePost)

	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w, r)
	case "POST":
		handlePost(w, r)
	case "PUT":
		handlePut(w, r)
	case "DELETE":
		handleDelete(w, r)
	}

}

func handleGet(w http.ResponseWriter, r *http.Request){
	return
}

func handlePost(w http.ResponseWriter, r *http.Request){
	return
}

func handlePut(w http.ResponseWriter, r *http.Request){
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request){
	return
}
