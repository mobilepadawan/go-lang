//Guide tutorial: https://www.youtube.com/watch?v=XZaBie7O0Rk
package main

//importamos la librería NET/HTTP para crear nuestro servidor
//importamos el paquete de datos fmt, que sirve para formatear datos en Go
//importamos el paquete strconv que sirve para convertir cadenas de string
import (
	"fmt"
	"net/http"
	"strconv"
)

//Creamos nuestra función Main (inicial) para levantar nuestro servidor web en Go
func main() {
	fs := http.FileServer(http.Dir("./static/"))              //definimos el fichero fileServer y le indicamos que, el directorio de los archivos estáticos es la carpeta STATIC. Todo queda guardado en la variable fs, correspondiente a file system
	http.Handle("/static/", http.StripPrefix("/static/", fs)) //Luego, con handle le decimos al servidor web que, cuando lleguen peticiones de contenido estático, sirva las mismas a través del prefijo indicado stripPrefix()

	http.HandleFunc("/", home)             //respuesta para home (una función Go apropiada con contenido HTML)
	http.HandleFunc("/info", info)         //respuesta para info (una función Go apropiada con información del servidor)
	http.HandleFunc("/producto", producto) //respuesta para producto (una función Go apropiada con info del total de productos)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	html := "<html>"
	html += "<body>"
	html += "<h1 class='hola'>Hola, mundo!</h1>"
	html += "</body>"
	html += "</html>"
	w.Write([]byte(html))
}

func info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Host: ", req.Host)
	fmt.Fprintln(w, "URI: ", req.RequestURI)
	fmt.Fprintln(w, "Method: ", req.Method)
	fmt.Fprintln(w, "RemoteAddr: ", req.RemoteAddr)
}

var productos []string //defino una variable productos, que será un array de strings

func producto(w http.ResponseWriter, req *http.Request) {
	html := "<html>"
	html += "<body>"
	html += "<h1 class='hola'>Productos</h1>"
	html += "<p class='hola'>Total de productos: " + strconv.Itoa(len(productos)) + "</p>"
	html += "</body>"
	html += "</html>"
	w.Write([]byte(html))
}
