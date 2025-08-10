/* - 1. Calculo de interes Simple - Crea un programa que calcule el interes simple usando la formula: `interes = capital * tasa * tiempo`:
 *- Usa constantes para la tasa de interes (5%).
 *- Usa variables para el capital ($1000) y el tiempo en años (3).
 *- Calcula y muestra el interes total.
 */

package ejercicios

import "fmt"

const tasaInteres = 0.05

func CalcularInteresSimple() {
	capital := 1000.0
	tiempo := 3
	interes := capital * tasaInteres * float64(tiempo)
	total := capital + interes
	fmt.Printf("Interes: $%.2f, Total: $%.2f\n", interes, total)
}
