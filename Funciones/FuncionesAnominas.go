package Funciones

import (
	"fmt"
)

func Calculos() {
	var numero3 int = 32
	var numero4 int = 243

	calculo := func(numero1 int, numero2 int) int { // Es una funcion anonima
		return numero1 + numero2 + numero3 + numero4
	}
	fmt.Println(calculo(10, 25))

	// Redefiniendo la funcion "calculo", se puede porque es una funcion anomin
	calculo = func(numero1 int, numero2 int) int { // Ya no se ocupa asignarla, solo iguala el valor.
		return numero1 * numero2 * numero3
	}
	fmt.Println(calculo(10, 25))

}
