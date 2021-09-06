package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//Funciones
func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una funcion"
}

//Handlers

func Index(rw http.ResponseWriter, r *http.Request) {
	funciones := template.FuncMap{
		"saludar": Saludar,
	}

	template, _ := template.New("index.html").
		Funcs(funciones).ParseFiles("index.html")
	template.Execute(rw, nil)
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
