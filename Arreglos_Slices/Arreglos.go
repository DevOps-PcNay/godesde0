package Arreglos_Slices

import (
	"fmt"
)

var tabla2 []int  // Coleccion de diferentes tipos de datos. Este es un Slides ya que no tiene definido el tamano.
var tabla [10]int // Es un arreglo, Vector, por defecto le asigna el valor 0 es el m pequeno, se inicia desde la posicion 0

// Asignando de otra manera
// No se requiere que se asignen todos
var tabla3 [20]int = [20]int{10, 0, 59}

// Usando cadenas de textos
var tabla4 [10]string = [10]string{"Juan", "Pedro", "Jose"}

// Declarando un arreglo bidireccional, Matriz
var matriz [20][30]int

func MuestraArreglos() {
	// Asignacion directa
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)
	fmt.Println(tabla3)
	fmt.Println(tabla4)

	// Recorriendo un arreglo
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	// Asignando valores a la Matriz
	matriz[7][24] = 15
	fmt.Println(matriz)
}
