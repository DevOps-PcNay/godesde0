package ejercicios

import (
	"strconv"
)

func Funcion(cadena string) (int, string) {
	// Esta funcion retorna dos valores, para omitr un valor :

	var cadena2 string
	numeros, err := strconv.Atoi(cadena)

	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}

	if numeros > 100 {
		cadena2 = "Es mayor a 100"
	}
	if numeros < 100 {
		cadena2 = "Es menor a 100"
	}

	return numeros, cadena2
}
