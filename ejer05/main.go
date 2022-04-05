package main

import "fmt"

func main() {
	//Primera forma:
	// i := 1
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//Forma 2
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	//Forma 3
	// numero := 0
	// for {
	// 	fmt.Println("Continuo")
	// 	fmt.Println("Ingrese número secreto: ")
	// 	fmt.Scanln(&numero)
	// 	if numero == 100 {
	// 		break
	// 	}
	// }

	//Forma 4
	// var i = 0
	// for i < 10 {
	// 	fmt.Printf("\n Valor de i: %d", i)
	// 	if i == 5 {
	// 		fmt.Printf(" multiplicamos por 2 \n")
	// 		i = i * 2
	// 		continue
	// 	}

	// 	fmt.Printf("       Paso por aqui \n")
	// 	i++
	// }

	//Forma 5
	var i int = 0

RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a RUTINA sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}
