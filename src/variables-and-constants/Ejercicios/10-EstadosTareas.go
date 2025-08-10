/* - 10. Estados de Tareas - Gestiona estados de tareas con `iota`:
 * - Define estados: Pendiente, EnProgreso, Completada.
 * - Crea una variable para el estado actual
 * - Cambia el estado de Pendiente a EnProceso
 * - Muestra el valor numerico en cada estado
 */

package ejercicios

import "fmt"

const (
	pendiente = iota
	enProgreso
	completada
)

func EstadosTareas() {
	estadoActual := pendiente
	fmt.Println("Estado actual:", estadoActual)
	estadoActual = enProgreso
	fmt.Println("Estado cambiado a:", estadoActual)
}
