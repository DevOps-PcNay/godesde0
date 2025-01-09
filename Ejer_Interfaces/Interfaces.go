package Ejer_Interfaces

import (
	"fmt"
	"godesde0/Interfaz"
)

// hu = Es una variable de tipo Interfaz.Humano, en este caso recibe "Hombre", "Mujer" cuando se llame la funcion
func HumanosRespirando(hu Interfaz.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n ", hu.Sexo())
}
