package logger

import (
	"errors"
	"testing"
)

func TestBuildAndLog(t *testing.T) {

	t.Run("Imprimir un string", func(t *testing.T) {

		texto := "Dos veces no es pa tanto xD"

		err := BuildAndLog(texto, "logger.TestBuildAndLog.t.Run")
		if err.Error() == "" {
			t.Fatal("El error está vacío")
		}
	})

	t.Run("Imprimir un error", func(t *testing.T) {

		err := errors.New("Tipo error")
		err = BuildAndLog(err, "logger.TestBuildAndLog.t.Run")
		if err.Error() == "" {
			t.Fatal("El error está vacío")
		}
	})
}
