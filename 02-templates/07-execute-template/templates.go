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

var templates = template.Must(template.New("T").ParseFiles("index.html", "base.html"))

//Handlers
func Index(rw http.ResponseWriter, r *http.Request) {
	usuario := Usuario{"Alex", 26}
	err := templates.ExecuteTemplate(rw, "index.html", usuario)

	if err != nil {
		panic(err)
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
