package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Configurar el directorio de sonidos y el puerto
	const multimedia = "recursos"
	const port = 8080

	//Controlador
	http.Handle("/", addHeaders(http.FileServer(http.Dir(multimedia))))
	fmt.Printf("Iniciando servidor en el puerto %v\n", port)
	log.Printf("El servicio %s se ejecuta en el puerto HTTP: %v\n", multimedia, port)

	//Control de errores
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

//addHeaders
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
