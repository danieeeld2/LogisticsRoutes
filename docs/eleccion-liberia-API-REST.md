# Framework API REST

## Criterios de elección

Los criterios que vamos a seguir para la elección de la herramienta son los siguientes:

- Que tenga un buen rendimiento, sea rápido y eficiente en el uso de recursos.
- Debido a las dimensiones de nuestro proyecto (que son pequeñas, ya que se trata de un proyecto para una asignatura de la universidad), buscamos una herramienta ligera, con las funcionalidades suficientes para nuestro objetivo de diseñar la API  REST.
- Facilidad para usar y configurar. En este punto incluimos que tenga un formato fácil de usar, por ejemplo `JSON` u otros que venimos comentando en objetivos anteriores.
- Compatibilidad y salud de los paquetes: Uno de los problemas principales que me he encontrado a lo largo del proyecto ha sido la compatibilidad de paquetes, así que lo tendremos en cuenta como un criterio importante (Recordemos que tenemos configurado [Snyk.io](https://snyk.io/) para verificar la salud de los mismos).

## Frameworks estudiados

### [Gin](https://gin-gonic.com/)

Framework minimalista enfocado en la velocidad y rendimiento. Permite el uso de `JSON`. Tiene las herramientas que necesitamos, es ligera y parece que se integra bien con repositorios de terceros.

### [Echo](https://github.com/labstack/echo)

Lo voy a descartar porque advierte de que puede tener incompatibilidades con repositorios de tercero, y dado los problemas que hemos tenido en los objetivos pasados, prefiero curar en salud.

### [Librería Estándar net/http](https://pkg.go.dev/net/http)

Teniendo en cuenta los criterios, es una elección bastante sólida, ya que no vamos a tener problemas de compatibilidad y de uso, al ser parte de la librería estándar.

### [Chi](https://go-chi.io/#/)

Es ligero y con buen rendimiento y se integra bien con la librería estándar `net/http`. Además. parece que tiene muy buena compatibilidad con repositorio de terceros.

### [Buffalo](https://gobuffalo.io/es/)

Me parece muy buena herramienta, pero es más robusta y ligera que las otras, así que, siendo consecuentes con los criterios que hemos elegido, la descarto.

### [Gorilla Mux](https://github.com/gorilla/mux)

Es una opción más potente, que permite configuraciones avanzadas de rutas mas avanzadas que,por ejemplo `net/http`, pero también es más complejo y robusto, así que, siguiendo los criterios, la descarto.

### [Revel](https://revel.github.io/)

Es una de las que más curiosas me han parecido. Cuenta con diseño MVC integrado, Hot Code Reload, gestión integrada de dependencias..., pero, nuevamente, incluye mucho más de lo que requerimos, así que la descartamos.

