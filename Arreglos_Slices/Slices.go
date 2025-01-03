package Arreglos_Slices

import (
	"fmt"
)

// Slices = Es un arreglo que no se definen cuantos elementos.

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78}

func MuestroSlices() {
	fmt.Println(tablaSlice)
	porcion := arreglo[3:]   // Slices creado desde un arreglo, contenido desde posicion 3
	porcion2 := arreglo[:5]  // Slices desde posicion 0 hasta 5
	porcion3 := arreglo[6:7] // Slices desde posicion 6 hasta 7
	fmt.Println("Porcion ", porcion)
	fmt.Println("Porcion2 ", porcion2)
	fmt.Println("Porcion3 ", porcion3)
}

func Capacidad() {
	// 5 Elementos y 20 de capacidad(espacio en memoria )
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n ", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d \n ", len(nums), cap(nums))

}
