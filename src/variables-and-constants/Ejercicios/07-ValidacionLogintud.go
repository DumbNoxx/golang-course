/* - 7. Validacion de Longitud - Verifica si una contraseña cumple con la longitud minima:
 * - Define una constante `loginMin = 8`
 * - Usa una variable `password` con un valor
 * - Determina si es valida comparando su longitud con `logMin` - Nota puedes usar la funcion `len(value)` que retorna un integer y el operador de comparacion `==`.
 * - Muestra "Valida" o "Invalida".
 */

package ejercicios

import "fmt"

const loginMin = 8

func ValidacionDeLongitud() {
	password := "1234567"
	result := len(password) == loginMin
	fmt.Println("El resultado es:", result)
}
