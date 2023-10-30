# Gestor de tareas

Necesitamos elegir un **Task Runner** para automatizar tareas repetitivas y ayudar en el proceso de compilación y ejecución de nuestro proyecto. Para ello, vamos a estudiar algunas de las opciones más interesantes con las que me he encontrado:

- **Make**: Es una de las herramientas más utilizadas para la automatización de tareas. Utiliza un fichero `Makefile` para definir el conjunto de reglas que permiten la ejecución de estas tareas. Tiene la ventaja de que es muy genérico, ya que se basa en el uso de la shell para la ejecución de tareas, permitiendo ser usado por gran variedad de lenguajes; lo que lo convierte en una opción muy versátil.

- **Mage**: Es una herramienta de construcción y automatización escrita en Go, la cual, permite definir tareas en este lenguaje; lo que lo hace que sea fácil de usar y extender a proyectos en Go. Puedes definir las tareas en el archivo `magefile.go` y ejecutarlas desde la línea de comando.

- **Task**: Es otra herramienta para la automatización de tareas en Go. Permite definir tareas en un archivo `Taskfile.yml` y ejecutarlas desde línea de comando.

Trás investigar un poco más, las dos opciones más sólidas son **Make** o **Mage**, así que veamos sus ventajas y desventajas:

#### Mage
##### Ventajas:

- Escrito en Go, por lo que es muy práctico para proyectos diseñados en este lenguaje.

- Alto grado de mantenibilidad, pudiendo definir tareas en un archivo `magefile.go` o,incluso, en varios y son fácilmente customizables para los distintos sistemas operativos.

- No tiene dependencias más allá del propio lenguaje Go.

- Trabajamos con el mismo lenguaje en el que programamos el proyecto.

##### Desventajas:

- Necesitas saber Go, pero como nuestro proyecto está en este lenguaje, podríamos descartar esta desventaja.

- No es tan versátil en otros lenguajes fuera de Go.

#### Make
##### Ventajas:

- Ámpliamente conocido y utilizado en el mundo del software.

- Multiplataforma y compatible con varios sistemas operativos y lenguajes.

##### Desventajas:

- Sintaxis confusa dependiendo del proyecto, lo que puede llevar a una menor claridad y mantenibilidad.

- Menos modularidad que Mage

Teniendo en cuenta los criterios reflejados en [Elección de ToolChain](https://github.com/danieeeld2/LogisticsRoutes/issues/15), **Mage** es una elección sólida para nuestro proyecto, el cuál, esta diseñado en Go; por lo que podemos definir directamente las tareas en este lenguaje, beneficiando así la mantenibilidad e integración natural del sistema. Por lo que, a pesar de no ser la opción más popular, esta vez nos quedaremos con **Mage** por las razones previamente explicadas. 

