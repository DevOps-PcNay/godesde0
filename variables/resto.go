package variables

import (
	"fmt"
	"strconv"
	"time"
)

// Variables que se declare fuera la funcion, esta sera visible en todas las funciones que se definen en este archivo "resto.go" e incluso fuera del paquete varaibles, pero debe iniciar con la letra mayuscula

// Funcion publica, letra primera
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	// Cuando la variable se definde dentro de la funcion solo es vista dentro de esta funcion.

	// Esta funcion es visible en los archivos que se encuentren en el paquete (carpeta variable )
	MuestroEnteros()
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

// Funcion que devuelve un valor.
// Segundo parentesis es donde retorno el valor.
func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
