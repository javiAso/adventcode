package getInput

import (
	"io/ioutil"
	"log"
	"net/http"
)

const COOKIE = "_ga=GA1.2.627316507.1671029597; _gid=GA1.2.1242918966.1671029597; session=53616c7465645f5fb614ab027226f11b9fea6d56f08c3297011530ce7fc11a2f3ac2b7a5a7d506b259633519054c5857fc5f0340d23b11f55b2cd79bf3023dd7"

func GetInput(url string) (string, error) {
	clienteHttp := &http.Client{}                     //Creamos el cliente HTTP
	peticion, err := http.NewRequest("GET", url, nil) // Creamos la peticion con 3 parametros: tipo, url y body (como esta es get el body va a nil)
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
		return "", err
	}
	// Podemos agregar encabezados
	peticion.Header.Add("Cookie", COOKIE)
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
