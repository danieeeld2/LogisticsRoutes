# Filosofía seguida a la hora de desarrollar los tests

La filosofía que vamos a seguir para desarrollar los test es **F.I.R.S.T**:

- **Fast**: Los tests deben ejecutarse de forma rápida. De esta forma, no tiene importancia ejecutar los test muchas veces, ya que se tarda poco tiempo en ver los resultados.

- **Isolated**: Los tests ejecutan partes aisladas denuestro código. En general, la mayoría de tests se encargan de probar partes aisladas del código y no dependen de datos externos o APIs.

- **Repeatable**: Los tests siempre devuelven los valores esperados y se testea acorde a estos.

- **Self-Validating**: Los datos de salida y entrada esperados son automáticos y no hace falta esperar ninguna acción manual.

- **Thorough**: Hay que proponer desde casos fáciles hasta complejos, para ver cómo se comporta el programa a partir de distintos datos de entrada.
