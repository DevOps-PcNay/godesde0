package Funciones

import (
	"fmt"
)

//Funcion Recursiva, se llama a si misma "n" veces

func Exponencia(valor int) {

	// Esta condicion siempre se escribe para que no se quede ciclada de forma indefinida
	if valor > 100 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2)

}
