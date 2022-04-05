package main

import "fmt"

//Funciones anónimas
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Restamos 6 - 4 = %d \n", Calculo(6, 4))

	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Dividimos 12 - 3 = %d \n", Calculo(12, 3))

	Operaciones()

	//Closures
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla()
	)
	}
}

func Operaciones() {

	//Funciones anónimas
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}

//Función que recibe un parámetro y como resultado me retorna una nueva función
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}

}
