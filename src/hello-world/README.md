# Primer "Hello, World" en Go

Bienvenido al primer programa que vas a escribir en Go, y en la que la mayoría de personas hacen en sus inicios de cualquier lenguaje de programación

Lo primero que tendríamos que hacer (más adelante te explico una mejor manera de hacerlo) es crear un archivo `go.mod` con este contenido


```mod
module Hello_World

go 1.24.4
```

Ya con ese archivo creado podemos empezar a programar. Crea un archivo llamado `main.go` con el siguiente contenido:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World")
}
```

Y felicidades, haz escrito tu primer programa en Go, ahora vamos a desglosar lo que escribimos.

```go
package main
```

Esta línea le indica a Go que este archivo pertenece al paquete especial llamado `main` siendo el punto de entrada para la ejecución del programa.

```go
import "fmt"
```

Esta línea indica a Go que estás importando el paquete `fmt` de la librería estándar de golang.

```go
func main(){
  
}
```

La declaración de la función main, es la función principal que le indica a Go que todo lo que este dentro de la función debe ejecutarse.

```go
fmt.Println("Hello, World")
```

Y esta línea indica que estás haciendo la función `Println` del paquete `fmt`, indicando que quieres imprimir en pantalla con un salto de línea al final de la
llamada a la función.

Ahora podés ejecutar esto en tu consola:

```bash
go run main.go

# Output: Hello, World
```

## Felicidades ya has escrito tu primer Hola mundo en Go, ya con esto podemos pasar al siguiente parte del curso: [aquí](../variables-and-constants/README.md)


## Pequeño plus (Esto lo explicaré más adelante pero lo vamos a usar a partir de acá)

Como dije anteriormente esto lo explicaré más adelante pero justamente esto lo vamos a usar a partir de ahora, para seguir con el curso.

Go viene incorporado con una serie de comandos para ejecutar, uno de ellos es para inicializar un proyecto con Go, el cual es:

```bash
go mod init github.com/nombreUsuario/RepositorioDelProyecto
```

Con ese comando se crea automáticamente el archivo `go.mod` vamos a desglosarlo.

```bash
go mod init
```

Le indica a Go que debe inicializar un archivo `.mod` donde después del init debería ir el nombre del proyecto, por convención siempre ponemos la
url del repositorio de **github**, mod init debe de tener un identificador único para el módulo o proyecto que tengas entonces usamos lo ante mencionado
aprovechando que los repositorios de github son únicos.

Y con esto terminamos este primer capítulo del curso de Go.
