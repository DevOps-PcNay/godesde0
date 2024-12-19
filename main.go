package main

import (
	// Export GOPATH = /var/www/html/soporte.my-empresa.com/github.com/DevOps-PcNay/godesde0
	"fmt"
	"godesde0/variables"
)

//"fmt",

func main() {
	//fmt.Println("Hola mundo")
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(13)
	fmt.Println(estado)
	fmt.Println(texto)

}
