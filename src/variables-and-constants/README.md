# Variables y Constantes

Las variables son las herramientas que usamos a diario los programadores, son las que almacenan los valores de algún tipo de dato.

En Golang se pueden declarar de dos maneras:

```Golang
var variable
variable1 :=
```

La primera se usa la palabra reservada (keyword) `var` y en la segunda se usa el operador de walrus el cual también se le dice el operador de asignación corta.

Las cual pueden quedar asi:

```Golang
var variable = "Esto es una variable"
variable1 := "Esto tambien es una variable"
```
```
```

En Golang también existen las Constantes, son un tipo de variable que almacenan valores inmutable, o que no pueden ser modificados, a diferencia de los demás lenguajes de programación, las constantes no tienen que ir en mayúsculas, a menos que quiera que sea exportada (Se hablará de ese tema más adelante). Al principio se usará el camelCase si la constante es local, o usando SCREAMING_CASE si es una constante de ámbito global, también puedes usar PascalCase para declarar constantes.


Para declarar una constante se tiene que usar la keyword `const`.

```Golang
const PORT_SERVER = 8080 // Constante de ambito global
const PortServer = 8080 // Variante de para una constante de ambito global o exportable
const apiKey = "https://example.com/api/auth/registration" // Constante de ambito local 
```


Ahora, la pregunta es, cuándo usar `var` o el `:=`, bien, esa pregunta es frecuente, pero con una respuesta sencilla al principio.

Var se usa más que todo para declarar una variable pero no inicializarla, siempre para cuando necesites que sea de un tipo implícito o que la inicialización ocurra en otro momento.

```Golang
var name string 
```
```
```

Que tambien puede ser inicializada sin ningun problema.

```Golang
var name string = "Dylan"
```

Pero para esos casos seria mejor usar el `walrus` para declarar e inicializar.


```Golang
name := "Dylan"
```

El compilador de Go se considera generalmente inteligente, debido a varias razones que no mencionaré acá, pero una de ella es que infiere los tipos de las variables en tiempo de compilación, aunque sea un lenguaje de tipado estático. Aunque no se declare explícitamente el tipo de la variable el compilador lo determina automáticamente y lo asocia permanentemente a la variable.


Por defecto cada tipo de dato a ser declarada se inicializa con el valor cero, esto significa que al declarar una variable con `var` y que sea `string` se inicializa con el valor `""`, ejemplo:

```Golang
var cadena string // Valor por default ""
var entero int // Valor por default 0
var flotante float64 // Valor por default 0
var booleano bool // Valor por default false
```


Golang es tan simple en sintaxis que incluso se te permite declarar múltiples variables como en Python, JavaScript o C#.

```Golang
var (
  name string
  age int
  active bool
)
```


De asimismo también te permite declarar constantes con la misma sintaxis.

```Golang
const (
  PORT_SERVER = 8080
  PortServer = 8080
  apiKey = "https://example.com/api/auth/registration" 
)
```

Mejorando la legibilidad en el código y agrupando un solo llamado a la keyword `const` o `var`.

Ahora hay una funcionalidad bastante peculiar en las declaraciones de constantes. Existe un identificador predeclarado integrado que simplifica las definiciones de constantes incrementables, quiero decir que representa un entero que se autoincrementa en cada constante de un bloque de constantes. 

```Golang
const (
  MONDAY = iota // 0
  TUESDAY // 1
  WEDNESDAY // 2
  THURSDAY // 3
  FRIDAY // 4
  SATURDAY // 5
  SUNDAY // 6
)
```


Especialmente es útil para definir un conjunto de constantes que comparten un tema o secuencia común. Evita la asignación manual de los valores de cada constante, lo que reduce la probabilidad de errores y hace el código más conciso. Ten en cuenta que también puedes hacer cálculos con iota, ejemplo:

```Golang
const (
  one = iota + 1 // 1
  two // 2
  three // 3
)
```

Ahora vendría el alcance y las sombras o comúnmente llamado (scope & shadowing). El `Scope` determina la accesibilidad de la variable en el código a nivel de bloque y el `Shadowing` esto ocurre cuando las variables de alcance interno ocultan a las variables de ámbito global siendo del mismo nombre. Golang tiene alcances de paquete, función y bloque, para comprender esto habría mirado el siguiente ejemplo.

```Golang
package main

import "fmt"

var dev = "Dylan" // Esta variable es de alcance en todo el archivo digo todo el archivo porque no es una variable exportable - Package Scope

func changeName(name string) string {
  var newName string
  dev = name // Hace shadowing a la variable de ambito global 
  newName = dev // Se utiliza el Functional Scope 
  return newName
}

func main(){
  fmt.Println(dev) // Output: Dylan

  newName := changeName("Kevin")
  fmt.Println(newName) // Output: "kevin"
  fmt.Pritln(dev) // Output: "Kevin"

  {
    dev = "Nazareth" // Vuelve a hacer shadowing a la variable de ambito global 
    blockScope = dev  // utiliza el Block Scope 
    fmt.Println(blockScope) // Output: Nazareth
  }
  fmt.Println(blockScope) // Error: Undefined blockScope
  fmt.Println(dev) // Output: Nazareth
}
```

No se enfoquen en la función `changeName` Luego veremos funciones a más profundidad. Lo importante es que se presten especial atención al `Scope & Shadowing`, se hace entender que si no tienes conocimientos previos sobre la programación, suele resultar confuso al principio pero con la práctica te volverás un maestro, puedes realizar los siguientes ejercicios para practicar todo lo aprendido en este apartado.


### Apartado de ejercicios

- 1. Cálculo de interés Simple - Crea un programa que calcule el interés simple usando la fórmula: `interes = capital * tasa * tiempo`:
  - Usa constantes para la tasa de interés (5%).
  - Usa variables para el capital ($1000) y el tiempo en años (3).
  - Calcula y muestra el interés total.

- 2. Sistema de niveles de usuario - Define 3 niveles de usuario usando `iota`: Básico, Intermedio, Avanzado:
  - Asigna el nivel "Intermedio" a una variable.
  - Muestra un mensaje con el nivel numérico y el nombre correspondiente.

- 3. Conversor de Unidades (Longitud) - Convierte 5 metros a centímetros y pulgadas: 
  - 1 metro = 100 cm.
  - 1 metro = 39.37 pulgadas.
  - Usa constantes para los factores de conversión.
  - Muestra ambos resultados

- 4. Gestión de Inventario Básico - Administra el inventario de un producto:
  - Declara variables para: nombre ("Camiseta"), stock (50), precio (25.99).
  - Actualiza el stock a 45 despues de una venta
  - Muestra los valores antes y despues de la actualizacion

- 5. Secuencia Aritmética - Genera una secuencia numérica con iota:
  - Crea constantes para 4 valores.
  - El primer valor debe de ser 10.
  - Cada siguiente valor debe incrementarse en 5.
  - Calcula y muestra la suma de todos los valores.

- 6. Shadowing de tipos - Demuestra shadowing con diferentes tipos:
  - Declara una variable global `valor (int = 100)`.
  - Muestra la variable global.
  - Crea una variable local usando `{}` -  `valor (string = "cien")`.
  - Cambia el valor de la variable global dentro de los corchetes
  - Muestra todos los valores fuera de los corchetes
 

- 7. Validación de Longitud - Verifica si una contraseña cumple con la longitud mínima:
  - Define una constante `loginMin = 8`
  - Usa una variable `password` con un valor
  - Determina si es válida comparando su longitud con `logMin` - Nota puedes usar la función `len(value)` que retorna un integer y el operador de comparación `==`.
  - Muestra "Válida" o "Inválida".

- 8. Niveles de Log - Crea un sistema de niveles de Log:
  - Usa `iota` para definir: Debug, Info, Warning, Error.
  - Asigna el nivel "Warning" a una variable.
  - Muestra el valor numerico y el nombre del nivel.

- 9. Cálculo de Área - Calcula el área de un rectángulo:
  - Declara variables para base (10.5) y altura (5.0).
  - Calcula el area (base * altura).
  - Muestra los resultado con dos decimales.

- 10. Estados de Tareas - Gestiona estados de tareas con `iota`:
  - Define estados: Pendiente, EnProgreso, Completada.
  - Crea una variable para el estado actual
  - Cambia el estado de Pendiente a En Progreso
  - Muestra el valor numérico en cada estado

Bien, con eso terminamos con la parte teórica, ahora espero se esfuercen en la parte práctica, mucha suerte y si te atascas en la carpeta que ejercicios están los ejercicios resueltos de esta lección, mucha suerte Gopher.

--- 

Si culminaste con esta lección por favor ve a la [Siguiente lección](../data-types/README.md)
