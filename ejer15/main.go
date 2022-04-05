package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//GoRoutines es el equivalente a las promesas
	//El go ejecuta de forma asincrona el método miNombreLento
	go miNombreLentooooo("Cristian Collazos")
	fmt.Println("Estoy aqui")

	var x string
	fmt.Scanln(&x)
}

func miNombreLentooooo(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)

	}
}
