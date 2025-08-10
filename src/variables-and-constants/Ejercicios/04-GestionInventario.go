/*- 4. Gestion de Inventario Basico - Administra el inventario de un producto:
 * - Declara variables para: nombre ("Camiseta"), stock (50), precio (25.99).
 * - Actualiza el stock a 45 despues de una venta
 * - Muestra los valores antes y despues de la actualizacion
 */

package ejercicios

import "fmt"

func GestionInventarioBasico() {
	nombre := "Camiseta"
	stock := 50
	precio := 25.99
	stock = 45
	fmt.Println("Valores luego de vender 5 productos:")
	fmt.Printf("Articulo: %s\nStock: %d\nPrecio: %.2f\n", nombre, stock, precio)
}
