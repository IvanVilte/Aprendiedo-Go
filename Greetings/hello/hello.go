package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	//Establece las propiedades del Logger ('registrado') predefinido, incluyendo el prefijo
	//de la entrada del log y una bandera para deshabilitar la impresión de la hora, el archivo fuente y el número de línea.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Un slice de nombres.
	names := []string{"Gladys", "Samantha", "Darrin"}

	//Solicitar un mensaje de bienvenida para los nombres.
	messages, err := greetings.Hellos(names)

	//Si se devuelve un error, imprimirlo en la consola y salir del programa.
	if err != nil {
		log.Fatal(err)
	}

	//Si no se devuelve ningún error, imprime el map de mensajes devuelto en la consola.
	fmt.Println(messages)
}
