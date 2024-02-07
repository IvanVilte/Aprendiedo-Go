package greetings

import "fmt"

// Hola devuelve un saludo para la persona nombrada.
func Hello(name string) string{
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}