# go-error-logger

Es un módulo que ayuda a tirar mis logs cuando se cachan errores, la idea es
pasarle el error y una frace para identificar el problema en consola, esto es
con la intención de poder dejar a una goruntina que ejecute el log sin que
me detenga o atrace la ejecución de mis programas.

## Uso

```
import "github.com/kalmecak/go-error-logger"


// Ejemplo 1
if err != nil {

    go logger.Error(err, "mipaquete.funcx.funcy")
}

// imprime:

// 2021-03-12T18:33:25.705797107Z: *** mipaquete.funcx.funcy ***
// 2021-03-12T18:33:25.705797107Z: Aquí va el mensaje del error
// 2021-03-12T18:33:25.705797107Z: --- mipaquete.funcx.funcy ---

// Ejemplo 2

err := errors.New("Error loco que destruye el planeta")

go logger.Error(err, "path/al/error")

// imprime:

// 2021-03-12T18:33:25.705797107Z: *** path/al/error ***
// 2021-03-12T18:33:25.705797107Z: Error loco que destruye el planeta
// 2021-03-12T18:33:25.705797107Z: --- path/al/error ---


// Ejemplo 3

err := "No es un error pero igual quiero un mensaje aquí"

go logger.Message(err, "path/al/mensaje")

// imprime:

// 2021-03-12T18:33:25.705797107Z: *** path/al/error ***
// 2021-03-12T18:33:25.705797107Z: No es un error pero igual quiero un mensaje aquí
// 2021-03-12T18:33:25.705797107Z: --- path/al/error ---

```