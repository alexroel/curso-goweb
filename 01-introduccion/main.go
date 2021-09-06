package main

import (
	"fmt"
	"log"
	"net/http"
)

//Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	//GET POST PUT DELETE
	//Obterner el metodo
	fmt.Println("El metodo es +" + r.Method)
	fmt.Fprintln(rw, "Hola Mundo")
}

//Pagina no encontrada
func PaginaNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

//Errores
func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Este es un Error", http.StatusConflict)
}

//Argumentos por URL
func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery) //Separa de url
	fmt.Println(r.URL.Query())  //Mapa

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(rw, "Hola %s tu edad es %s!!", name, age)

}

func main() {
	//Mux
	mux := http.NewServeMux()

	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/pagina", PaginaNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	//Router
	/*
		http.HandleFunc("/", Hola)
		http.HandleFunc("/pagina", PaginaNF)
		http.HandleFunc("/error", Error)
		http.HandleFunc("/saludar", Saludar)
	*/

	//Crear Servidor
	//http.ListenAndServe("localhost:3000", nil)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	log.Fatal(server.ListenAndServe())
	//log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
