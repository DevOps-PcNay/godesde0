package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Va a devolver un valor, de tipo "string"
func TablaMultiplicar() string {
	var (
		numero1      int
		numero_texto string
		err          error
		texto        string
	)

	for a := 0; a < 10; a++ {
		teclado := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingresa un numero ")

		if teclado.Scan() {
			numero1, err = strconv.Atoi(teclado.Text())
			if err != nil {
				fmt.Println("El dato es incorrecto ")
				continue

			} else { // if err != nil
				break
			} //  // if err != nil

		} //if teclado.Scan() {

	} //for a := 0; a > 10; a++ {

	numero_texto = strconv.Itoa(numero1)
	fmt.Println("EL dato es correcto  \n " + numero_texto)

	// Imprimiendo la tabla. 1 al 10 del numero seleccionado.
	for b := 1; b <= 10; b++ {
		if b == 1 {
			texto += fmt.Sprintf("\n")
		}
		texto += fmt.Sprintf("%d x %d = %d \n ", numero1, b, b*numero1)

	} //for b := 1; b <= 10; b++ {

	return texto

} // func TablaMultiplicar {
