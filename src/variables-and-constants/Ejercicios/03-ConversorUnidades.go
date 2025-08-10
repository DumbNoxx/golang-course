/* - 3. Conversor de Unidades (Longitud) - Convierte 5 metros a centimetros y pulgadas:
 * - 1 metro = 100 cm.
 * - 1 metro = 39.37 pulgadas.
 * - Usa constantes para los factores de conversion.
 * - Muestra ambos resultados
 */

package ejercicios

import "fmt"

const (
	conversionCentimetros = 100.0
	conversorPulgadas     = 2.54
)

func ConversorUnidades() {
	metro := 1
	centimetros := metro * conversionCentimetros
	fmt.Printf("1 metro en centimetros es: %dcm\n", centimetros)
	pulgadas := (conversionCentimetros / conversorPulgadas) * float64(metro)
	fmt.Printf("1 metro en pulgadas es %.2fP\n", pulgadas)
}
