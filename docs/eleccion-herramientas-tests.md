# Elección de metodología y herramientas para test

Podemos ver una idea general de todo lo que vamos a plantear en [(M1-Dev) Elección de las herramientas de test y de la metodología #26](https://github.com/danieeeld2/LogisticsRoutes/issues/26)

## Testing en Go

En **Go**, el test runner más utilizado es el nativo del propio lenguaje, ya que cumple los criterios planteados en [(M1-Dev) Elección de las herramientas de test y de la metodología #26](https://github.com/danieeeld2/LogisticsRoutes/issues/26) y, al ser parte del propio lenguaje, cuenta con el propio soporte del mismo y es un estándar de Go. Algunas características de cómo se testea de forma nativa en Go son las siguientes:

- En vez de aserciones, tiene errores que tienes que especifican si alguno de los resultados no ha sido el esperado. Además, como Go no cuenta con excepciones, habrá que controlar los errores devolviendo valores específicos, que nos permitirán comprobar el correcto funcionamiento de nuestro código.

- Permite ejecutar benchmarks (Más información [Aquí](https://pkg.go.dev/testing#hdr-Benchmarks)) para medir el rendimiento del código.

- Permite implementar mocks mediante el uso de interfaces, creando [funciones tipadas](https://go.dev/doc/effective_go#functions)

- También tenemos la posibilidad de crear y ejecutar subsets o, incluso, sub-benchmarks. Podemos obtener más información acerca de esto pinchando [aquí](https://pkg.go.dev/testing#hdr-Subtests_and_Sub_benchmarks)


- También, mirando el apartado de [Fuzzing](https://pkg.go.dev/testing#hdr-Fuzzing) de la documentación anterior, vemos que nos permite generar inputs aleatorios con el fin de encontrar bugs en nuestro código, para casos que no habíamos anticipado en los tests.

Como podemos observar, las funcionalidades de testing que nos dá nativamente Go son extraordinarias. Aun así, hay librerías de aserciones y paquetes de Go que también pueden resultar útil y que hacen la vida más fácil. Veamos algunos casos:

### Goblin 

Se trata de un framework de testeo que sigue la filosofía BDD. Tiene la desventaja de que no incluye tantas funcionalidades como otros paquetes que veremos ahora, pero sí tiene la curiosidad de que permite describir el propósito de cada test mediante strings y tiene una salida por consola bastante vistosa, que ayuda a inspeccionar los sucesos ocurridos durante los tests.

Como la mayoría de paquetes en Go, podemos encontrar la información en su [repositorio](https://github.com/franela/goblin)

### Testify

Se trata de una colección de paquetes que permite realizar [aserciones](https://github.com/stretchr/testify#assert-package), [mocking](https://github.com/stretchr/testify#require-package), así como ejecutar código antes y después de los conjuntos de tests que hayamos diseñado (A esta técnica se le conoce como [suite](https://github.com/stretchr/testify#suite-package)). Toda la documentación referente la podemos encontrar en el propio [repositorio del proyecto](https://github.com/stretchr/testify).

Una funcionalidad curiosa con la que me he topado leyendo la documentación, es la posibilidad de combinar las aserciones con [require](https://github.com/stretchr/testify#require-package), que permite terminar los test que fallen las aserciones.

Junto a las opciones nativas de Go, es una de las opciones más populares gracias al buen soporte que ha tenido hasta día de hoy y la gran cantidad de funcionalidades que incorpora.

### GoCheck

A nivel de funcionalidad, es bastante similiar a **Testify**, tal y como podemos consultar en su [repositorio](https://labix.org/gocheck), aunque su sintaxis no es tan intuitiva. Una característica curiosa es la posibilidad de crear y gestionar [directorios temporales](https://pkg.go.dev/gopkg.in/check.v1#C.MkDir)

#### Gomega

Está principalmente dirigido a la creación de aserciones y es muy útil para realizar pruebas siguiendo la filosofía BDD, sin embargo, tal y como podemos observar en su [repositorio](github.com/onsi/gomega), no es tan útil como las otras opciones, al centrarse simplemente en las aserciones.


## Decisión

Sin duda, contamos con una gran variedad de opciones. No obstante, teniendo en cuenta nuestros criterios, las dos más interesantes serían la biblioteca nativa y testify. 

Por un lado, **Testify** es potencialmente más complejo y requiere de una dependencia externa, pero, a cambio, proporciona aserciones más expresivas. Por otro lado, la **Biblioteca de Pruebas Estándar de Go** es la opción más utilizada y nos libera de la necesidad de usar dependencias, además de que funciona con el comando `go test`, pero, sus aserciones son más básicas y sus funcionalidades algo más limitadas; por ejemplo, no tenemos la posibilidad de crear mocks directamente como en Testify, aunque, como bien hemos dicho antes, lo podemos diseñar nosotros mismos mediante el uso de interfaces con funciones tipadas.

Como vemos es una elección difícil y cualquiera de las dos es perfectamente válida. En mi caso y, teniendo en cuenta, tal y como digo en #26, por simplicidad y, ya que se trata de un proyecto para una asignatura, voy a elegir la **Librería Estándar** por ser más simple, no incluir más dependencias y, lo más importante, por ser el estándar del lenguaje. Pese a esto la otra opción es totalmente válida y se podría elegir sin problemas como hemos visto.


