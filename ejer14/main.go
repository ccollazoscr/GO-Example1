package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//ejemploDefer()
	ejemploPanic()
}

func ejemploDefer() {
	//Ejemplo de defer
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	//Esta función no se ejecuta secuencialmente. Se ejecuta cuando salga de la función
	defer f.Close()

	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1)
	}
}

func ejemploPanic() {
	//Esta función se lanza cuando se lanze un panic dentro de un método
	//Defer ejecuta una sola cosa. Si quiero hacer varias... debo crear una fx anónima
	defer func() {
		//Recover me trae el resultado del último panic. Si no hubo panic, reco va a ser nulo
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		//Panic aborta el programa en GO
		panic("Se encontro el valor de 1")
	}
}
