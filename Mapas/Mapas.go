package Mapas

import (
	"fmt"
)

func MostrarMapas() {
	// Definiendo un mapa
	paises := make(map[string]string)
	paises["Mexico"] = "Cd. mexico"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)           // Mostrando todo la estrucutura del mapa
	fmt.Println(paises["Mexico"]) // Mostrando por Clave, nos muestra

	// Otra forma de definir el mapa
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"chivas":       37,
		"Boca Juniors": 30,
	}
	// Se puede ir agregando no tiene instruccion "append" ya que no tiene limite establecido al crearlo.

	fmt.Println(campeonato)

	// Recorriendo el mapa
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n ", equipo, puntaje)
	}

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	// Para mostrar un elemento, muestra uno de los "valores", otro valor para este caso "existe"
	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe = %t \n", puntaje, existe)

}
