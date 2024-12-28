package Files

import (
	"bufio"
	"fmt"
	"godesde0/ejercicios"
	//"io/ioutil"
	//"ioutil"
	"os"
)

var fileName string = "Files/Txt/Tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(fileName)
	// Crea el arhivo, si existe o borra y lo crea de nuevo
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return // Sale de este metodo.
	}

	// Graba en el archivo.
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(fileName, texto) { //if Append(fileName, texto) == false {
		fmt.Println("Erroral Concatenar Contenido")

	}
}

func Append(filen string, texto string) bool {
	// Abre el archivo para agregar contenido
	// 0664 = Permiso de escritura al Owner(dueno) y demas Grupo, Others es solo Lectura.
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println("Error Durnte Agregando el archivo " + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error al Grabar el archivo " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.ReadFile(fileName)
	// ioutil.Readfile = Esta depreciada desde la version
	if err != nil {
		fmt.Println("Error Al Leer Archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo)) // Lo convierte de bytes a Cadena
}

// Leer archivo usando "Open"
func LeoArchivo2() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("hubo un error Al Leer Archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(" > " + registro)
	}
	archivo.Close()

}
