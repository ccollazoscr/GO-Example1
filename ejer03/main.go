package main

import "fmt"

var estado bool

func main() {
	estado = true
	//No son ncesarios los parentesis si no se necesitan
	//En la misma linea se puede asignar un valor
	if estado = false; estado == true {
		fmt.Println(estado)
	} else if estado == false {
		fmt.Println("Es False")
	}

	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")
	}
}
