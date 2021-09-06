package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Estructuras
type Usuario struct {
	UserName string
	Edad     int
}

//Handlers
func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("index.html")

	usuario := Usuario{"Alex", 26}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, usuario)
	}
}

//Función principal
func main() {

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	//Server
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}
