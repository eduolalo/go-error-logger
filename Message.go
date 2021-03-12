package logger

import "log"

// Message imprime el string en consola usando la locación para que sea fácil
// determinar dónde se envió
func Message(err, location string) {

	log.Println("*** ", location, " ***")
	log.Println(err)
	log.Println("--- ", location, " ---")
}
