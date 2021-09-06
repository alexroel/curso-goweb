package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//Estructuras
type Usuario struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

//Handlers

func Index(rw http.ResponseWriter, r *http.Request) {

	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"JavaScript"}

	cursos := []Curso{c1, c2, c3, c4}
	usuario := Usuario{"Alex", 26, true, false, cursos}
	template, err := template.ParseFiles("index.html")
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
