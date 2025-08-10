/* - 8. Niveles de Log - Crea un sistema de niveles de Log:
 * - Usa `iota` para definir: Debug, Info, Warning, Error.
 * - Asigna el nivel "Warning" a una variable.
 * - Muestra el valor numerico y el nombre del nivel.
 */

package ejercicios

import "fmt"

const (
	debug = iota
	info
	warning
	error
)

func NivelesLog() {
	nivel := warning
	fmt.Printf("Nivel: %d, Nombre: Warning\n", nivel)
}
