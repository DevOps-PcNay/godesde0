package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	var (
		numero1        int
		multiplicacion int
		numero_texto   string
		numero3_texto  string
		err            error
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
	fmt.Println("EL dato es correcto " + numero_texto)

	// Imprimiendo la tabla. 1 al 10 del numero seleccionado.
	for b := 1; b <= 10; b++ {
		numero2_texto := strconv.Itoa(b)
		multiplicacion = numero1 * b
		numero3_texto = strconv.Itoa(multiplicacion)

		fmt.Println("Numero : " + numero_texto + " X " + numero2_texto + " = " + numero3_texto)

		// Ahora usando fmt con formato para numeros

		fmt.Println("Usando fmt con texto formateado para numeros y letras mixtos ")
		fmt.Printf("%d x %d = %d \n ", numero1, b, b*numero1)

	} //for b := 1; b <= 10; b++ {

} // func TablaMultiplicar {
