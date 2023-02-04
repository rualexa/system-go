package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Inicio)
	log.Println("servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
