package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(2)

	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println("Calculo: ")
	fmt.Println(Calculo(10, 30))
	fmt.Println(Calculo(20, 30))
	fmt.Println(Calculo(30, 30))
}

//Opción 1
// func uno(numero int) int {
// 	return numero * 2
// }

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

//Función con parámetros variables. Se van a recibir n números enteros
func Calculo(numero ...int) int {
	total := 0
	//Range itera por cada elemento de número.
	//Range devuelve dos valores. El primer valor que devuelve es el número del elemento. Si no se va a usar el parámetro se puede utilizar se asigna _
	//for _, num := range numero {
	for item, num := range numero {
		total = total + num
		fmt.Printf("\nItem %d \n", item)
	}

	return total
}
