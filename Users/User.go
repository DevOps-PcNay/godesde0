package Users

import (
	"fmt"
	"godesde0/Modelos"
	"time"
)

func AltaUsuario() {
	u := new(Modelos.User) // Se crea el objeto, antes de "new" era una definicion.
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
