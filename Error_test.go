package logger

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {

	t.Run("Imprimir un error", func(t *testing.T) {

		err := errors.New("Un errorsillo los comete cualquiera")

		Error(err, "logger.TestError.t.Run")
	})
}
