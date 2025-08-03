# Primer "Hello, World" en Go

Bienvenido al primer programa que vas a escribir en Go, y en la que la mayoria de personas hacen en sus inicios de cualquier lenguaje de programacion

Lo primero que tendriamos que hacer (mas adelante te explico una mejor manera de hacerlo) es crear un archivo `go.mod` con este contenido


```mod
module Hello_World

go 1.24.4
```

Ya con ese archivo creado podemos empezar a programar. Crea un archivo llamado `main.go` con el siquiente contenido:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World")
}
```

Y felicidades, haz escrito tu primer programa en Go, ahora vamos a desloglar lo que escribimos.

```go
package main
```

Esta linea le indica a Go que este archivo pertenece al paquete especial llamado `main` siendo el punto de entrada para la ejecucion del programa.

```go
import "fmt"
```

Esta linea indica a Go que estas importando el paquete `fmt` de la libreria standard que viene incorporada en golang.

```go
func main(){
  
}
```

La declaracion de la funcion main, es la funcion principal que le indica a Go que todo lo que este dentro de la funcion debe ejecutarse.

```go
fmt.Println("Hello, World")
```

Y esta linea indica que estas haciendo la funcion `Println` del paquete `fmt`, indicando que quieres imprimir en pantalla con un salto de linea al final de la
llamada a la funcion.

Ahora podes ejecutar esto en tu consola:

```bash
go run main.go

# Output: Hello, World
```

## Felicidades ya has escrito tu primer Hola mundo en Go, ya con esto podemos pasar al siquiente parte del curso: [aqui](../variablesAndConstants/README.md)


## Pequeño plus (Esto lo explicare mas adelante pero lo vamos a usar apartir de aca)

Como dije anterior mente esto lo explicare mas adelante pero justamente esto lo vamos a usar apartir de ahora, para seguir con el curso.

Go viene incorporado con una serie de comando para ejecutar, uno de ellos es para inicializar un projecto con Go, el cual es:

```bash
go mod init github.com/nombreUsuario/RepositorioDelProyecto
```

Con ese comando se crea automaticamente el archivo `go.mod` vamos a desglosarlo.

```bash
go mod init
```

Le indica a Go que debe inicializar un archivo `.mod` donde despues del init deberia ir el nombre del proyecto, por convencion siempre ponemos la
url del repositorio de **github**, mod init debe de tener un identificador unico para el modulo o proyecto que tengas entonces usamos lo ante mencionado
aprovechando que los repositorios de github son unicos.

Y con esto terminamos este primer capitulo del curso de Go.
