package Defer_panic

import (
	"fmt"
	"log"
	//"os"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Ete es el segundo mensaje ")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero Panic \n %v ", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
