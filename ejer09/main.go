package main

import "fmt"

func main() {
	//Mapas son parecidos a los diccionarios
	//Mapas: Conjunto de elementos que se pueden serializar
	//Se utiliza para reservar espacios. Go adiciona mas en caso de que se adicionen más de 5
	//paises := make(map[string]string,5)
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	fmt.Println(campeonato)

	//Asignar valores
	campeonato["River Plate"] = 25
	//Actualizar valores
	campeonato["Chibas"] = 55
	fmt.Println(campeonato)

	//Borrado
	delete(campeonato, "Real madrid")
	fmt.Println(campeonato)

	//El range lo entrega en desorden. Siempre que se ejecuta lo presenta de forma diferente
	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d\n", equipo, puntaje)
	}

	//Verificar si existe o no en un map
	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje es  %d,y el equipo es: %t\n", puntaje, existe)
}
