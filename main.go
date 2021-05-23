package main

import (
	"log"

	"github.com/juanhayek/goarg/bd"
	"github.com/juanhayek/goarg/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a base de datos")
		return
	}
	handlers.Manejadores()
}
