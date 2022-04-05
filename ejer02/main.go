package main

import (
	"fmt"
	"strconv"
)

//Las variables que inician con minuscula.. son privadas
//Si inician con mayuscula son globales. Aplica apra los métodos también.
var numero int
var texto string
var status bool

func main() {
	//var numero2, numero3, numero4, numero5 int
	//El := me crea una nueva variable
	numero2, numero3, numero4, texto, status, texto2 := 2, 3, 4, "Esto es una prueba:", true, ""

	var moneda int64 = 0
	numero2 = int(moneda) //Forma de castear tipos

	//Forma de convertir un int a string
	texto = fmt.Sprintf("%d", moneda)

	texto2 = strconv.Itoa(int(moneda))

	numero5 := 10
	numero5 = 15 + 4

	fmt.Println(texto)
	fmt.Println(texto2)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)
	fmt.Println(status)

	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
