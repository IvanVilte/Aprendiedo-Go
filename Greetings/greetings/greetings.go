package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)		

// 'Hola' devuelve un saludo para la persona nombrada.
func Hello(name string) (string, error){
	
	// Si no se ha dado ningún nombre, devuelve un error con un mensaje.
	if name == "" {
		return "", errors.New("Empty name.")
	}
	
	//Crea un mensaje utilizando un formato aleatorio.	
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

//randomFormat devuelve uno de un conjunto de mensajes de saludo. El mensaje devuelto se selecciona al azar.
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