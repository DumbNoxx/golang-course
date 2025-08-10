/*- 2. Sistema de niveles de usuario - Define 3 niveles de usuario usando `iota`: Basico, Intermedio, Avanzado:
 *- Asigna el nivel "Intermedio" a una variable.
 *- Muestra un mensaje con el nivel numerico y el nombre correspondiente.
 */

package ejercicios

import "fmt"

const (
	basico = iota
	intermedio
	avanzado
)

func SistemaNiveles() {
	valorIntermedio := intermedio
	fmt.Println("Basico: ", basico)
	fmt.Println("Intermedio: ", intermedio)
	fmt.Println("Avanzado: ", avanzado)
	fmt.Println("Valor asignado a a la variable:", valorIntermedio)

}
