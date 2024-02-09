package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main()  {
	
	//Establece las propiedades del Logger ('registrado') predefinido, incluyendo el prefijo
	//de la entrada del log y una bandera para deshabilitar la impresión de la hora, el archivo fuente y el número de línea.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Solicitar un mensaje de bienvenida.
	message, err := greetings.Hello("Gladys")
	
	//Si se devuelve un error, imprimirlo en la consola y salir del programa.
	if err != nil {
		log.Fatal(err)
	}

	//Si no se devuelve ningún error, imprima el mensaje devuelto en la consola.
	fmt.Println(message)
}