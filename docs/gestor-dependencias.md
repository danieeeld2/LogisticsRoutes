# Gestor de dependencias

Desde las versiones posteriores a `Go 1.11` (del año 2018) se introdujo el concepto de módulos a este lenguaje, que son, colecciones de paquetes relacionados que se versionan juntos, proporcionando todas las dependencias necesarias para su base de código, a través de la creación de **Go Module**. 

No existe un gestor de dependencias como tal, sino que las dependencias se gestionan a través de **Go Module**, a través de los siguientes archivos:
- `go.mod`: Archivo donde se definen los module path de un módulo. Este archivo contiene metadatos sobre un módulo, como su nombre y las dependencias que requiere.
- `go.sum`: Archivo encargado de gestionar las versiones e indicar la integridad de las mismas.

Además, la gestión de dependencias en este lenguaje es descentralizada, lo que permite tratar un enlace a un repositorio git como un módulo más.

Por último, **Go Module** incluye una serie de comandos que nos ayudan con la gestión de estas dependencias, como son: `go get`, `go mod tidy`,..., entre otros.
