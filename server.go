package main

import (
	"log"
	"net/http"
	"tryMaria/routes"
)

func main() {
	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", routes.M))
}
