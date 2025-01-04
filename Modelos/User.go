package Modelos

import (
	"time"
)

// Por convecciones de Go Land, en el archivo "user.go" se definen los "types" y "estructuras"
type User struct {
	// Entidad "User"
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// "this" hace referencia al objeto en cuestion, para este caso es "user"
// (this User) = Es la parte fundmental
// *User = Apunta a una direccion de memoria.

func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAt
	usuario.Status = status
}
