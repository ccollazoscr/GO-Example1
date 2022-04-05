package main

import (
	"fmt"
	"time"

	us "./user"
)

//Ejemplo 1
// type usuario struct {
// 	Id        int
// 	Nombre    string
// 	FechaAlta time.Time
// 	Status    bool
// }

//Ejemplo 2 con herencia
type pepe struct {
	us.Usuario
}

func main() {
	//Ejemplo 1
	// User := new(usuario)
	// User.Id = 10
	// User.Nombre = "Pablo"
	// //GO no tiene clases. Utiliza estructuras.
	// fmt.Println(User)

	u := new(pepe)
	u.AltaUsuario(1, "Cristo", time.Now(), true)
	fmt.Println(u.Usuario)
}
