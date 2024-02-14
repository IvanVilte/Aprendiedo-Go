package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// 'Hola' devuelve un saludo para la persona nombrada.
func Hello(name string) (string, error) {

	// Si no se ha dado ningún nombre, devuelve un error con un mensaje.
	if name == "" {
		return "", errors.New("Empty name.")
	}

	//Crea un mensaje utilizando un formato aleatorio.
	//message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat())

	return message, nil
}

// Hellos devuelve un mapa que asocia a cada una de las personas nombradas un mensaje de saludo.
func Hellos(names []string) (map[string]string, error) {

	//Un map para asociar nombres con mensajes.
	messages := make(map[string]string)

	//Recorre en bucle el slice de nombres recibido, llamando a la función Hello para obtener un mensaje para cada nombre.
	for _, name := range names {

		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//En el map, asocia el mensaje recuperado con el nombre.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat devuelve uno de un conjunto de mensajes de saludo. El mensaje devuelto se selecciona al azar.
func randomFormat() string {
	//Un conjunto de formatos de mensaje.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	//Devuelve un formato de mensaje seleccionado aleatoriamente especificando un índice aleatorio para la rebanada de formatos.
	return formats[rand.Intn(len(formats))]
}
