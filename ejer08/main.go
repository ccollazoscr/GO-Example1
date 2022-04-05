package main

import "fmt"

//Ejemplo 1
// var tabla [10]int

// func main() {
// 	tabla[0] = 1
// 	tabla[5] = 15

// 	fmt.Println(tabla)
// }

//Ejemplo 2
// func main() {
// 	tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	for i := 0; i < len(tabla); i++ {
// 		fmt.Println(tabla[i])
// 	}

// }

//Ejemplo 3
// var matriz [5][7]int

// func main() {
// 	matriz[3][5] = 1
// 	fmt.Println(matriz)
// }

func main() {
	//Si no se le asigna longitud.... se puede asignar los valores que se quieran
	//Se le conoce como slice
	matriz := []int{2, 5, 4}
	fmt.Println(matriz)

	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	//porción crea un slice teniendo a partir de la 3ra posición de elementos
	porcion := elementos[3:]

	fmt.Println(porcion)
}

func variante3() {
	//slice con 5 elementos. El 20 se utiliza para reservar espacio en memoria.
	//Inicialmente solo utiliza 5 elementos
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
