package main

import (
	"fmt"
	"time"
)

func main() {
	//GoRoutines es el equivalente a las promesas
	//El go ejecuta de forma asincrona el método miNombreLento

	//Chanels son unos canales que permiten que una go rutine envien información hacia otra función.
	canal1 := make(chan time.Duration)
	go bulce(canal1)
	fmt.Println("Llegué hasta aquí")

	//Hasta que no haya un valor en canal1 no se continua con la aplicación
	//Esto es parecido a un await
	msg := <-canal1
	fmt.Println(msg)
}

func bulce(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000; i++ {

	}

	final := time.Now()
	canal1 <- final.Sub(inicio)
}
