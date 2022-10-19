package logger

import (
	"errors"
	"testing"
)

func TestBuild(t *testing.T) {

	t.Run("Imprimir un string", func(t *testing.T) {

		texto := "Dos veces no es pa tanto xD"

		err := Build(texto, "logger.TestBuildAndLog.t.Run")
		txt := err.Error()
		if txt == "" {
			t.Error("El error está vacío")
		}
		if txt == "Dos veces no es pa tanto xD" {
			t.Error("El error no es el esperado, se obtuvo: ", txt)
		}
	})

	t.Run("Imprimir un error", func(t *testing.T) {

		err := errors.New("Tipo error")
		err = Build(err, "logger.TestBuildAndLog.t.Run")
		txt := err.Error()
		if txt == "" {
			t.Error("El error está vacío")
		}
		if txt == "Tipo error" {
			t.Error("El error no es el esperado, se obtuvo: ", txt)
		}
	})
}
