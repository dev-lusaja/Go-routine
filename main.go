package main

// Se utiliza el metodo Mensaje para ejecutarlo con 2 goroutines y una llamada directa,
// las 2 goroutines van a crear un nuevo hilo para ese metodo.

import (
	"fmt"
	"time"
)

func Mensaje(msj string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msj)
		time.Sleep(100)
	}
}

func main() {
	go Mensaje("con rutina")
	go Mensaje("con rutina2")
	Mensaje("sin rutina")
}
