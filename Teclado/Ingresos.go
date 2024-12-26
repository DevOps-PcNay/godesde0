package Teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

// No recibe ni envia numeros, por lo tanto es un metodo
func IngresoNumeros() {

	// Lee informacion del teclado, este paquete contiene mucha funciones para manejo de archivos.
	// os.Stdin = Entrada por teclado.
	// Es de tipo texto siempre.
	teclado := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1 : ")
	if teclado.Scan() { // Si teclado, es de tipo Cadena
		// No es necesario := por que se definio en "var numero1 int"
		numero1, err = strconv.Atoi(teclado.Text())
		if err != nil {
			// Se aborta el programa.
			panic("El dato es Incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2 : ")
	if teclado.Scan() { // Si teclado, es de tipo Cadena
		// No es necesario := por que se definio en "var numero2 int"
		numero2, err = strconv.Atoi(teclado.Text())
		if err != nil {
			// Se aborta el programa.
			panic("El dato es Incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingresa Leyenda  ")
	if teclado.Scan() { // Si teclado, es de tipo Cadena
		// No es necesario := por que se definio en "var numero2 int"
		leyenda = teclado.Text()
	}

	fmt.Println(leyenda, numero1*numero2)

} // func IngresoNumeros() {
