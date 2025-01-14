package main

// Export GOPATH = /var/www/html/soporte.my-empresa.com/github.com/DevOps-PcNay/godesde0

//"fmt"
//"godesde0/Teclado"
//"godesde0/ejercicios"
//"godesde0/variables"
//"runtime"
//"godesde0/Iteraciones"
//"godesde0/Files"
//"godesde0/Funciones"
// "godesde0/Arreglos_Slices"
//"godesde0/Mapas"
// 	"godesde0/Users"
import (
	//"godesde0/Ejer_Interfaces"
	//"godesde0/Modelos"
	//"godesde0/Defer_panic"
	//"fmt"
	//"godesde0/Goroutines"
	//"godesde0/WebServer"
	"godesde0/Middleware"
)

//"fmt",

func main() {
	/*
			//fmt.Println("Hola mundo")
			variables.MuestroEnteros()
			variables.RestoVariables()
			estado, texto := variables.ConviertoaTexto(13)
			fmt.Println(estado)
			fmt.Println(texto)


		variables.ConviertoaTexto(15)
		// os := runtime.GOOS // Determina que sistema Operativo

		if os := runtime.GOOS; os == "linux" || os == "OS X" {
			fmt.Println("El Sistema Operativo es Linux", os)
		} else {
			fmt.Println("El Sistema Operativo es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}

		variable1, variable2 := ejercicios.Funcion("2kjkjk000")
		fmt.Println("Variable 1 ", variable1)
		fmt.Println("Variable 2 ", variable2)
	*/

	//Teclado.IngresoNumeros()
	//Iteraciones.Iterar()

	// Pedir un valor desde el teclado
	//fmt.Println(ejercicios.TablaMultiplicar())
	//Files.GrabarTabla()
	//Files.SumaTabla()
	//Files.LeoArchivo()
	//Files.LeoArchivo2()
	//Funciones.Calculos()
	//Funciones.LlamarClosure()
	//Funciones.Exponencia(3)
	//Arreglos_Slices.MuestraArreglos()
	//Arreglos_Slices.MuestroSlices()
	//Arreglos_Slices.Capacidad()
	//Mapas.MostrarMapas()
	//Users.AltaUsuario()
	/*
		Pedro := new(Modelos.Hombre) // Esta instanciando el objeto, "Clase"
		Ejer_Interfaces.HumanosRespirando(Pedro)

		Maria := new(Modelos.Mujer)
		Ejer_Interfaces.HumanosRespirando(Maria)
	*/
	//Defer_panic.VemosDefer()
	//Defer_panic.EjemploPanic()

	// Goroutines.MiNombreLentoo("Ramon Ortega")

	// Para ejecutarlo de forma asincrona
	// Al ejecutarlo empieza y termina.
	// Esto pasa debido aque ejecuta esta linea (inicia pero como esta con sleep y primer milisegundo) y como no se tiene mas instrucciones termina.
	// Por lo que se tiene que tener cuidado la ejecucion de este programa.
	// Se utilizan mucho en operaciones muy pesadas que se ejecuten en 2do plano y ejecutando otras instruccions, pero se debe terner cuidado en verificar que se terminen la Go routines.

	/*
		canal1 := make(chan bool)
		go Goroutines.MiNombreLentoo("Ramon Ortega", canal1)

		// Esta esperando que se oprima una tecla, pero se esta ejecutado en el segundo plano en desplegar las letras de los nombres cada 1 seg., hasta que se oprima una tecla y Enter.

		// Al llegar a esta linea ya asigno el valor de "canal1 = true", ya termino el ciclo "for".

		// Ahora utilizando el "defer" con una funcion anomima
		defer func() {
			<-canal1 // <-canal1 // await de como se usa en NodeJs
			// Se pueden agregar mas canales y no termina la ejecucion del programa hasta que se complete los canales
		}()
		fmt.Println("Aun estoy ejecutando la funcion -MiNombreLentoo- ")
	*/
	// WebServer.MiWebServer()
	Middleware.MiMiddleware()

} // func main()
