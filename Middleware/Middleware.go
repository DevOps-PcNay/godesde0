package Middleware

import (
	"fmt"
)

func suma(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	// Se pasa como parametro una funcion
	result := operacionMiddleware(suma)(2, 3) // Solo se escribe la primera vez : :=
	fmt.Println("Resultado de la Suma = ", result)

	result = operacionMiddleware(restar)(10, 6)
	fmt.Println("Resultado de la Resta= ", result)

	result = operacionMiddleware(multiplicar)(2, 4)
	fmt.Println("Resultado de la Multiplicacion = ", result)

}

// Se define el Middleware se define con los mismo parametros que se definen en la funcion principal (origen) y retorna una funcion que recibe dos parametros Enteros y que retorno un valor Entero.
func operacionMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		// En esta parte se puede agregar las validaciones necesarias de los parametros de la funcion o del cuerpo de la funcion.
		fmt.Println("Inicio de Operacion ")
		// Para salir de la funcion se tiene que retornar "return"
		// Es lo que esta retornando la funcion "operacionMiddleware"
		return f(a, b)

		// Por lo que solo el middleware("operacionMiddleware") solo esta desplegando un pantalla un mensaje.

	}
}
