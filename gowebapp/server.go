package main

//importamos la librería NET/HTTP para crear nuestro servidor
import (
	"net/http"
)

//Creamos nuestra función Main (inicial) para levantar nuestro servidor web en Go
func main() {
	fs := http.FileServer(http.Dir("./static/"))             //definimos el fichero fileServer y le indicamos que, el directorio de los archivos estáticos es la carpeta STATIC. Todo queda guardado en la variable fs, correspondiente a file system
	http.Handle("/static", http.StripPrefix("/static/", fs)) //Luego, con handle le decimos al servidor web que, cuando lleguen peticiones de contenido estático, sirva las mismas a través del prefijo indicado stripPrefix()
	http.ListenAndServe(":8080", nil)
}
