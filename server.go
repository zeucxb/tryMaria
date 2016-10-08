package main

import (
	"log"
	"net/http"
	"tryMaria/config"
	"tryMaria/routes"
)

func main() {
	log.Println("Server is running...")
	log.Println(http.ListenAndServe(config.GetPort(), routes.M))
}
