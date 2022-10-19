package logger

import (
	"log"
)

// Error imprime el error en consola usando la locación para que sea fácil
// determinar dónde ocurrió
func Error(err error, location string) {

	log.Println("*** ", location, " ***")
	log.Println(err.Error())
	log.Println("--- ", location, " ---")
}
