package getInput

import (
	cookie "adventcode/cookie"
	"io/ioutil"
	"log"
	"net/http"
)

func GetInput(url string) (string, error) {
	clienteHttp := &http.Client{}                     //Creamos el cliente HTTP
	peticion, err := http.NewRequest("GET", url, nil) // Creamos la peticion con 3 parametros: tipo, url y body (como esta es get el body va a nil)
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
		return "", err
	}
	// Podemos agregar encabezados
	peticion.Header.Add("Cookie", cookie.COOKIE)
	respuesta, err := clienteHttp.Do(peticion) // Hacemos la peticion y guardamos la respuesta
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
		return "", err
	}
	defer respuesta.Body.Close()                           // Cerramos el cuerpo de la respuesta
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body) // Leemos el cuerpo de la respuesta
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
		return "", err
	}
	respuestaString := string(cuerpoRespuesta)
	/* 	log.Printf("Código de respuesta: %d", respuesta.StatusCode)
	   	log.Printf("Encabezados: '%q'", respuesta.Header)
	   	contentType := respuesta.Header.Get("Content-Type")
	   	log.Printf("El tipo de contenido: '%s'", contentType) */
	// Aquí puedes decodificar la respuesta si es un JSON, o convertirla a cadena
	//log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
	return respuestaString, nil

}
