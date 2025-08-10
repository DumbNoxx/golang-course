/* - 9. Calculo de Area - Calcula el area de un rectangulo:
 * - Declara variables para base (10.5) y altura (5.0).
 * - Calcula el area (base * altura).
 * - Muestra los resultado con dos decimales.
 */

package ejercicios

import "fmt"

func CalcularArea() {
	base, altura := 10.5, 5.0
	area := base * altura
	fmt.Printf("El area del rectangulo es: %.2f\n", area)
}
