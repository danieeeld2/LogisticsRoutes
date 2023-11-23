# LogisticsRoutes

###### Daniel Alconchel Vázquez

## Descripción del problema

Un problema fundamental para muchas empresas dedicadas al sector de la logística es la gestión eficiente de los vehículos de reparto.

## Propuesta de solución

Aplicación que permita optimizar las rutas de entrega, optimizándolas en función de distintos factores, como pueden ser el tráfico a tiempo real, las restricciones de peso y dimensión de los vehículos o la prioridad de las entregas, entre otros factores.

La aplicación persigue optimizar las rutas y reducir costos, permitiendo una mayor eficiencia y, por tanto, rentabilidad de la empresa.

## Documentación

- [Configuración Clave SSH](https://github.com/danieeeld2/LogisticsRoutes/blob/Objetivo-0/docs/ssh-key.png)

-  [Configuración Git](https://github.com/danieeeld2/LogisticsRoutes/blob/Objetivo-0/docs/gitconfig.png)

- [Personas](docs/personas.md)

- [Historias de Usario](docs/historias.md)

- [Milestones](docs/milestones.md)

- [Gestor de tareas](docs/gestor-tareas.md)

- [Gestor de dependencias](docs/gestor-dependencias.md)

- [Elección herramientas test](docs/eleccion-herramientas-tests.md)

- [Elección de imágen](docs/seleccion-imagen.md)

## Herramientas utilizadas

- **Lenguaje de programación**: Go

- **Task Manager**: Mage
	- `mage build` : Construye el proyecto
	- `mage installdeps` : Instalación de dependencias
	- `mage run` : Ejecuta el programa
	- `mage clean` : Limpia el proyecto
	- `mage check` : Comprueba la sintaxis
	- `mage test` : Realiza los test y te dá la salida por terminal
	
- Como podemos ver el el punto anterior, se ha añadido la tarea **check** para comprobar la sintaxis de las entidades.
- Como podemos ver hemos añadido también la tarea **test** para realizar los tests.

## Contenedor de pruebas

Se puede construir una imagen del contenedor y ejecutarla con

```bash
docker build -t danieeeld2/logisticsroutes . & docker run -t -v `pwd`:/app/test danieeeld2/logisticsroutes
```

También, podemos usar la imagen que está en [DockerHub](https://hub.docker.com/r/danieeeld2/logisticsroutes)

```bash
docker run -t -v `pwd`:/app/test danieeeld2/logisticsroutes:latest
```