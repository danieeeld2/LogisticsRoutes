# Sistema de Configuración

## Criterios

Los criterios que vamos a tener en cuenta para realizar la elección sobre la biblioteca de configuración son los siguientes:

- Facilidad de uso y legibilidad
- Formato de la configuración. Las bibliotecas pueden admitir diferentes formatos de configuración, como `JSON`, `YAML`... En mi caso, opto por usar bibliotecas que admitan formato `YAML`, por su facilidad y legibilidad (como detallo en el punto anterior).
- Soporte y seguridad de la biblioteca.
- Que nos permita manejar más de una configuración (Así podemos crear una configuración estándar para los tests y visible en git y otras privadas)
- Entensibilidad
- Salud de los paquetes, la cuál, servá verificada usando [Snyk.io](https://snyk.io/)


**Nota de procedimiento al diseñar la configuración**: Al utilizar bibliotecas de configuración, nos aseguramos de que la configuración esta separada del resto del código, pudiendo modificarse fácilmente, sin necesidad de cambiar todo el código.

## Opciones Estudiadas

### [Viper](https://pkg.go.dev/github.com/dvln/viper#section-readme)

Es una librería potente y flexible, permitiendo la carga de variables de entorno desde diferentes fuentes (como puede archivos `.env`, archivos de configuración `JSON`, `YAML`, ...), variables de entorno del sistema ...

Cuenta, también, con funciones adicionales para acceder a las variables de configuración de manera sencilla.

### [Godotenv](https://pkg.go.dev/github.com/Valgard/godotenv)

Se trata de una librería simple y fácil de usar, que se enfoca en la carga de variables de entorno desde archivos `.env`. Además, permite cargar múltiples archivos, así como especificar rutas personalizadas. Tiene la desventaja de que solo soporta archivos `.env`.

### [etcd/clientv3](https://pkg.go.dev/go.etcd.io/etcd/client/v3)

Se enfoca en el almacenamiento de claves-valor para mantener la configuración y el estado del sistema. Proporciona una interfaz para interactuar con el clusted de `etcd` y cuenta con varios métodos, que permiten realizar operaciones de transacciones, seguimiento de cambios, ...

### [Koanf](https://pkg.go.dev/github.com/knadh/koanf)

Es una alternativa basada en **Viper**, más moderna, ligera y con menos dependencias. Permite lectura de configuraciones de distintos ficheros, así como variables de entorno `.env` y permite usar otros formatos como `JSON`, `YAML`...

## Decisión final

Trás probar las opciones comentadas, me he encontrado con que **Viper** no consigo instalarlo correctamente en mi proyecto, por una incompatibilidad de las dependencias que no consigo solventar. El resto de opciones sí las he podido probar y he decidido quedarme con **Koanf**, por ser la más completa de las tres opciones anteriores, aunque, teniendo en cuenta las dimensiones de nuestro proyecto, cualquiera de las opciones era válida.