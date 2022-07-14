package main

import (
	"fmt"
	"genealog/rutas"
	"net/http"
)

func main() {
	fmt.Println("Iniciando la API")

	server := http.Server{

		Addr:    ":4000",
		Handler: rutas.Routing(),
	}
	fmt.Println("Api corriendo en el puerto 4000")

	server.ListenAndServe()

}
