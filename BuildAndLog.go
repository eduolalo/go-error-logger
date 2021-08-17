package logger

import (
	"errors"
	"log"
	"strings"
)

// BuildAndLog se encarga de recibir un error y una dirección de dónde se produjoe el error,
// regresando un error con el mismo string que se tira en el log
func BuildAndLog(err interface{}, location string) error {

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

	helper.WriteString("*** ")
	helper.WriteString(location)
	helper.WriteString(" ***\n")
	tolog := helper.String()
	log.Println(tolog)
	return errors.New(tolog)
}
