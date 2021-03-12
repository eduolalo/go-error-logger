package logger

import (
	"testing"
)

func TestMessage(t *testing.T) {

	t.Run("Imprimir un string", func(t *testing.T) {

		err := "Dos veces no es pa tanto xD"

		Message(err, "logger.TestMessage.t.Run")
	})
}
