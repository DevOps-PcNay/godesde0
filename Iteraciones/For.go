package Iteraciones

import (
	"fmt"
)

// Es un metodo, no recibe ni envia datos.
func Iterar() {
	var i int
	i = 0

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 5; j++ {
		fmt.Println("For simplificado ", j)
	}

	for k := 0; k <= 100; k += 20 {
		fmt.Println("For simplificado, con Increntos variados ", k)
	}

	for l := 100; l > 10; l -= 10 {
		fmt.Println("For simplificado, con Decremento variados ", l)
	}
	for m := 10; m > 1; m-- {
		fmt.Println("For simplificado, con Decremento uno por uno ", m)
	}

	for n := 10; n > 1; n-- {
		if n == 5 {
			break
		}

		fmt.Println("For simplificado, con Decremento uno por uno ", n)
	}

	for n := 10; n > 1; n-- {
		if n == 5 {
			continue
			// No mostrara el numero 5, y continura hasta 2
		}

		fmt.Println("For simplificado, con Decremento uno por uno ", n)
	}

}
