package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Usuarios"))
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5001", nil))
}
