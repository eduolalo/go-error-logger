package logger

import (
	"log"
)

// BuildAndLog se encarga de recibir un error y una dirección de dónde se produjoe el error,
// regresando un error con el mismo string que se tira en el log
func BuildAndLog(err interface{}, location string) error {

	toLog := Build(err, location)
	log.Println(toLog.Error())
	return toLog
}
