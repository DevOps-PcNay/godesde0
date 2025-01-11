package Goroutines

import (
	"fmt"
	"strings"
	"time"
)

// El programa no ejecuta otras instrucciones solo esta, por lo que no puede aprovechar la multitareas.
func MiNombreLentoo(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "") // Vector que se conforma por la separacion (espacio=) de la cadena
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) // 1 Segundo
		fmt.Println(letra)
	}

	// Va a disparar para que la Go Routine sepa que ya se tiene un valor para el fluj de ejecucion
	// Puede ser cualquier valor, pero para  este caso es de tipo "true"

	canal1 <- true
}
