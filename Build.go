package logger

import (
	"errors"
	"strings"
)

// Build se encarga de recibir un error y una dirección de dónde se produjoe el error,
// regresando un error con el mismo string que se tira en el log
// El err recibido puede ser de tipo string o de tipo error
// el error regresado es el que se ha construido, no un error de la ejecución de la función
func Build(err interface{}, location string) error {

	var helper strings.Builder
	helper.WriteString("*** ")
	helper.WriteString(location)
	helper.WriteString(" ***\n")
	switch t := err.(type) {
	case string:
		helper.WriteString(t)
	case error:
		helper.WriteString(t.Error())
	default:
		helper.WriteString("Ocurrio un error")
	}

	helper.WriteString("------ ")
	helper.WriteString(location)
	helper.WriteString(" ------\n")
	tolog := helper.String()
	return errors.New(tolog)
}
