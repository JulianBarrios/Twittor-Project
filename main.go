package main

import (
	"log"

	"github.com/JulianBarrios/Twittor-Project/bd"
	"github.com/JulianBarrios/Twittor-Project/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion con la db")
	}
	handlers.Handler()

}
