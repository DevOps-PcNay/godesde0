package WebServer

import (
	"net/http" // Para paquetes referentes a Servidor Web
)

func MiWebServer() {
	// Manajenndo las ruta
	http.HandleFunc("/", home)
	http.ListenAndServe(":3033", nil)
}

// Dos interfaz estandar que utiliza el paquete de "net/http"
// Revisar la ruta , se agrega "WebServer" ya que el "main.go" se encuentra en el directorio raiz, por lo que el archivo "index.html" se encuentra dentro de la carpeta "WebServer"
func home(Req http.ResponseWriter, Res *http.Request) {
	http.ServeFile(Req, Res, "./WebServer/index.html")
}
