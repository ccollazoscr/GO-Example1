package user

import "time"

//Se debe colocar en mayuscula para poder exportarlo
type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//This es un puntero al usuario
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
