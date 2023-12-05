# Sistema de logs

## Criterios

Los criterios que tendremos en cuenta para elegir una biblioteca de logs son:

- Soporte continuado en el tiempo
- Posibilidad de especificar distintos niveles de impresión: logs, debug, warning, error...
- Posibilidad de imprimir/almacenar los logs en diferentes ubicaciones: consola, ficheros locales, ficheros en un servidor remoto...

## Opciones Estudiadas

### [Log](https://pkg.go.dev/log)

Es parte de la librería estándar de Go, por lo que tiene la ventaja de que sabemos que va a tener un soporte continuado en el tiempo. Por otro lado, no tiene un gran conjunto de configuraciones, pero permite configurar:

- Nivel de registro: Permite establecer los distintos niveles de impresión, comentados anteriormente en los criterios (información, error, ...)
- Salida del registro: Permite imprimir los registros por la salida estándar, establecer un fichero de salida o una conexión de red para enviar los registros a través de ella.
- Prefixo de registro: Permite agregar un prefijo personalizado a los mensajes de registro, lo que es útil para identificar de qué parte de la aplicación provienen los registros.

El formato de salida de la librería incluye la fecha-hora, nivel de registro y el mensaje de registro, pero no tiene la facilidad de cambiar el formato de salida.

### [Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)

Se trata de una biblioteca de log estructurado para Go, que ofrece varias ventajas. Alguna de las que me han parecido más interesantes son las siguientes:

- Permite asociar clave-valor a los mensajes de log, facilitando en análisis y filtrado de los mismos.
- Admite diferentes niveles de log.
- Tiene una gran flexibilidad en el formato de salida, pudiendo elegir texto plano, JSON, ...
- Cuenta con gestión de hooks, permitiendo enviar los logs a varios destinos.
- Es compatible con el paquete `context` de Go, lo que permite proporcionar información contextual a través de las llamadas de funciones y asociarla con los logs generados.

### [Zerolog](https://pkg.go.dev/github.com/rs/zerolog)

Se trata de otra biblioteca de log estructurado para Go, que cuenta con ventajas muy similares a **Logrus**, además de una sintaxis simple.

## Decisión final

Para nuestro proyecto, las 3 opciones son válidas y sólidas. Voy a descartar la biblioteca estándar, porque me interesa tener un formato de salida más claro y personalizable para mi proyecto. Finalmente, entre **Logrus** y **Zerolog**, he decidido quedarme con **Zerolog** por que considero que es más fácil de configurar y la sintaxis es algo más simple, aunque cualquiera de las dos opciones hubiese sido válida.
