package Funciones

import (
	"fmt"
)

// Es un funcion locl solo se puede ver en este archivo(Clousers.go)
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	// Esta parte es el "Clouser"
	return func() int { // Funcion anonima
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	tabladel := 2
	MiTabla := tabla(tabladel) // Recibe el "Clouser"
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
