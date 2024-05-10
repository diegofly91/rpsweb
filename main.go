package main

import (
	"log"
	"net/http"
	"os"
	"rpsweb/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	log.SetPrefix("main: ")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// crear enrutador
	router := http.NewServeMux()

	// manejador de archivos estáticos
	fileStatic := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fileStatic))

	// rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := os.Getenv("PORT")
	log.Printf("Server running in port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
