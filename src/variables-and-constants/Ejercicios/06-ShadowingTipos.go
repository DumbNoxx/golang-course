/* - 6. Shadowing de tipos - Demuestra shadowing con diferentes tipos:
 * - Declara una variable global `valor (int = 100)`.
 * - Muestra la variable global.
 * - Crea una variable local usando `{}` -  `valor (string = "cien")`.
 * - Cambia el valor de la variable global dentro de los corchetes
 * - Muestra todos los valores fuera de los corchetes
 */

package ejercicios

import "fmt"

var integer int = 100

func ShadowingYTipos() {
	{
		valor := "cien"
		// El compilador te dice que esta variable no esta siendo utilizada
		fmt.Print(valor)
		integer = 200
	}
	valor := 9
	fmt.Println("Los valores son:", valor, integer) // Esto dara error porque "valor" esta fuera del scope y dentro del scope no se esta usando "valor"
	// Output: Undefined valor
}
