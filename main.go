package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mvitta/server/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /list", routes.HandleTestList)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
