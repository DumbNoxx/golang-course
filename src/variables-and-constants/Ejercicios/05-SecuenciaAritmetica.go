/* - 5. Secuencia Aritmetica - Generea una secuenta numerica con iota:
 * - Crea constantes para 4 valores.
 * - El primer valor debe de ser 10.
 * - Cada siquiente valor debe incrementarse en 5.
 * - Calcula y muestra la suma de todos los valores.
 */
package ejercicios

import "fmt"

const (
	one = 10 + (iota * 5)
	two
	three
	four
)

func SecuenciaAritmetica() {
	var suma int
	suma += one
	suma += two
	suma += three
	suma += four

	fmt.Println("Suma total de la secuencia numerica:", suma)
}
