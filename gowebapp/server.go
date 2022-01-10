//Guide tutorial: https://www.youtube.com/watch?v=XZaBie7O0Rk
package main

import (
	"fmt"		//importamos el paquete de datos fmt, que sirve para formatear datos en Go
	"net/http"	//importamos la librería NET/HTTP para crear nuestro servidor
	"strconv"	//importamos el paquete strconv que sirve para convertir cadenas de string
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

func producto(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	add, okForm := r.Form["add"] //peticiono el form, quien me devuelve 2 params: "add" si me devuelve algo. "okForm" si me devolvió algo o no (boolean)
	if okForm && len(add) == 1 {
		productos = append(productos, string(add[0])) //la función append() agrega el producto al array 'productos'
		w.Write([]byte("Producto agregado correctamente"))

		return
	}

	prod, ok := r.URL.Query()["prod"]
	if ok && len(prod) == 1 {
		pos, err := strconv.Atoi(prod[0])
		if err != nil {
			return
		}
		html := "<html>"
		html += "<body>"
		html += "<h1 class='hola'>Productos</h1>"
		html += "<p class='hola'>Producto: " + productos[pos] + "</p>"
		html += "</body>"
		html += "</html>"
		w.Write([]byte(html))
	}

	html := "<html>"
	html += "<body>"
	html += "<h1 class='hola'>Productos</h1>"
	html += "<p class='hola'>Total de productos: " + strconv.Itoa(len(productos)) + "</p>"
	html += "</body>"
	html += "</html>"
	w.Write([]byte(html))
}
