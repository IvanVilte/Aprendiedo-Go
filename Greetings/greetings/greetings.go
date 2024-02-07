package greetings

import (
	"errors"
	"fmt"
)		

// 'Hola' devuelve un saludo para la persona nombrada.
func Hello(name string) (string, error){
	
	// Si no se ha dado ningún nombre, devuelve un error con un mensaje.
	if name == "" {
		return "", errors.New("Empty name.")
	}
	
	// Si se ha recibido un nombre, devuelve un valor que incrusta el nombre en un mensaje de saludo.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}