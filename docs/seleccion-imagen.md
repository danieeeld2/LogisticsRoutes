# Elección Imagen

## Criterios

Los principales criterios a la hora de elegir una imagen son:

- Seguridad
- Rendimiento 
- Estándares
- Soporte

Teniendo esto en cuenta, la elección de imágenes antiguas quedan descartadas, ya que:

- A nivel de **seguridad** son inferiores, ya que las nuevas versiones suelen suplir errores de seguridad encontrados en algún momento.
- A nivel de **rendimiento**, (no siempre) el hecho de que sea una versión más reciente suele hacer que sea más rápida (en el contexto de construir la imágen), aunque hay que tener cuidado con el tamaño de las imágenes.
- A nivel de **estándares**, las nuevas funcionalidades de los lenguajes suelen ser útiles para seguir innovando y progresando en el desarrollo; por lo que tener versiones avanzadas (más recientes) permitesacar provecho de esas nuevas funcionalidades.
- A nivel de **soporte**, las imágenes más recientes tienen soporte más prolongado.

Además de todo esto, deberíamos tener en cuenta que la imagen solo incluya las herramientas necesarias para el proyecto.

## Consideraciones

### Debe tener Go ya instalado

Las imágenes de docker están diseñadas para ofrecernos lo mínimo imprescindible para poder hacer uso del software que queremos con el mínimo número de dependencias posibles. Esto puede implicar que la imagen oficial del lenguaje ofrezca un menor número de dependencias para ocupar menos contenedores.

De todas formas, esto no siempre es así y es importante estudiar bien las imágenes que se están planteando usar y cuáles opciones hay disponibles

## Imágenes candidatas

Tras investigar un poco, las principales imágenes que he encontrado, para este proyecto, son las siguientes:

- [Golang](https://hub.docker.com/_/golang)
- [Ubuntu](https://hub.docker.com/_/ubuntu)
- [Alpine](https://hub.docker.com/_/alpine)

### Alpine

Tiene la ventaja de que es muy liviana, pero no es tan sencilla de usar como Ubuntu y, además, debemos instalar y configurar Go nosotros mismos.

### Ubuntu

Tiene el problema de que es una imagen grande y que incluye más cosas de lo necesario. Además, también debemos instalar y configurar Go nosotros mismos.

### Golang

Nos ofrece una gran variedad de opciones, como, por ejemplo, combinar Golang con Alpine, lo cuál suple el problema que comentábamos anteriormente. Como su nombre indica, tiene la ventaja de que ya tiene Go instalado, lo cuál hace que sea más fácil de mantener nuestro Dockerfile. Además, al combinarlo con Alpine, tenemos una imágen pequeña con buen rendimiento.

Vamos a usar la versión [Golang 1.21](https://go.dev/doc/devel/release#go1.21.0) (ya que es la última versión, es estable y contiene las últimas mejoras en seguridad) con Alpine. Si entramos [Golang-Alpine](https://hub.docker.com/_/golang/tags?page=1&name=alpine) vemos que hay múltiples opciones. La más reciente es Golang 1.21 con Alpine 3.18, pero, tál y como vemos ahí, se especifica que hay más vunerabilidades; así que, siguiendo con nuestros criterios, usaremos [Golang 1.21-Alphine 3.17](https://hub.docker.com/layers/library/golang/1.21-alpine3.17/images/sha256-23b99eb736d97c576627bce305d93d14cf8074cd5da6a4e8b2f4767e9b381614?context=explore)

Toda la documentación necesaria sobre su uso están en [Golang Docker Official Image](https://hub.docker.com/_/golang)