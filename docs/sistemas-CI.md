# Sistemas de Integración Continua

## Criterios de selección

- Que sea gratis o que se pueda usar parte de su funcionalidad de forma gratuita, ya que, en nuestro caso, al ser un proyecto público pequeño, enfocado a una asignatura, no sería conveniente tener que pagar dinero por dichas funcionalidades.
- Que se puedan integrar con GitHub (use GitHub Check), ya que, a parte de que es necesario para la asignatura; es recomendable para que solo se hagan cambios y fusiones una vez pasado los tests.
- Que permita el uso de Go y Docker, que es la configuración que estamos usando en nuestro proyecto.

## Propuestas de sistemas

### [Jenkins](https://www.jenkins.io/)

Es una plataforma de código abierto que facilita la integración continua, así como el despliegue contínuo. Es altamente extensible mediante plugins y permite automatizar tareas mediante pipelines. Se puede configurar fácilmente con GitHub para que detecte automáticamente cambios en el repositorio y ejecute dichos pipelines. Además, es fácil de instalar y configurar mediante la interfaz web, siendo este servicio gratuito.

Pendiente de probar

### [Circle CI](https://circleci.com/)

Otra herramienta CI/CD que permite conectar GitHub. Esta herramienta permite la creación de trabajos individuales (jobs), secuencia de trabajos (pipelines) y definir como se ejecutan los pipelines (workflow). Admite el entorno de **Go y Docker** y es gratuito.

### [Semaphore](https://semaphoreci.com/)

Es otra opción que cumple con nuestros criterios, menos por el uso de GitHub Check. Esta opción hace uso de [Statuses](https://docs.github.com/es/rest/commits/statuses?apiVersion=2022-11-28). A pesar de no usar Github Check, se podría usar la opción proporcionada por Semaphore para prohibir megear un PR si alguno de los status contiene errores. Tampoco permite instalar gratuitamente en Github, por lo que queda descartada.

### [Azure Pipelines](https://azure.microsoft.com/es-es/products/devops)

He visto que algún compañero ha tenido problemas integrandolo, pero es una opción bastante famosa y, además, gestionada por Microsoft. Aun así, vamos a intentar probarla.

De momento, hemos creado el pipeline, pero para que se pueda ejecutar de forma gratuita, tenemos que hacer una petición que tarda de 2-3 días laborales a ser tramitada, por lo que lo dejaremos mientras en espera.