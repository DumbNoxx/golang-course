# Curso de Go siguiendo roadmap.sh

## Introducción a Go

### ¿Por qué usar Go?

#### ¿Qué es Go?

Go es un lenguaje de programación fácil de aprender, que combina la velocidad de C++ con la facilidad de sintaxis de Python. Fue creado en Google en 2007 y anunciado como proyecto de código abierto en 2012, con el lanzamiento de su versión 1.0.

A diferencia de otros lenguajes, Go fue diseñado para ser utilizado en los entornos distribuidos de Google. El propósito era crear un lenguaje que resolviera los grandes desafíos de la nube de Google, ofreciendo a los desarrolladores un entorno más eficiente y simple. En palabras de uno de sus creadores:

> "Go is more about software engineering than programming language research." - Rob Pike

Esto se traduce como: "Go trata más sobre ingeniería de software que sobre investigación de lenguajes de programación". Fue creado con el objetivo de resolver los siguientes puntos clave:

1.  **Compilaciones lentas:** Ofrecer tiempos de compilación rápidos.
2.  **Dependencias no controladas:** Mejorar la gestión de dependencias.
3.  **Lenguajes de subconjunto distintos (DSL):** Evitar la necesidad de lenguajes específicos de dominio.
4.  **Código ilegible:** Fomentar la legibilidad y la simplicidad.
5.  **Duplicación de esfuerzo:** Reducir la duplicación de código.
6.  **Actualizaciones costosas:** Facilitar las actualizaciones del lenguaje y las herramientas.
7.  **Incompatibilidad de versiones:** Minimizar los problemas de compatibilidad entre versiones.
8.  **Dificultad para escribir herramientas de automatización:** Simplificar la creación de herramientas.
9.  **Compilación en distintas plataformas:** Facilitar la compilación cruzada.

Los creadores de este lenguaje de programación son científicos de la computación de renombre:

*   **Ken Thompson:** Creador del lenguaje B, predecesor de C.
*   **Rob Pike:** Quien, junto a Ken Thompson, creó el sistema operativo Unix.
*   **Robert Griesemer:** Que trabajó en proyectos como el motor V8 de Google.

El resultado de la colaboración de estas tres mentes es Go.

#### Características de Go

*   **Lenguaje compilado:** El código se compila a binario nativo.
*   **Tipado estático y fuerte:** El tipo de las variables se comprueba en tiempo de compilación.
*   **Inferencia de tipos de datos:** El compilador puede deducir el tipo de las variables.
*   **Manejo de paquetes y módulos:** Sistema de organización de código.
*   **Amplia biblioteca estándar:** Incluye una gran cantidad de paquetes listos para usar.
*   **Tiempos de compilación rápidos:** Gracias a su análisis de dependencias y el uso de goroutines.
*   **Administración de memoria mediante punteros:** Permite un control de bajo nivel de la memoria.

Go no es un lenguaje orientado a objetos (OOP), ya que no tiene clases ni herencia, aunque puede presentar un paradigma similar en algunos aspectos. Aunque se pueden ver funciones anónimas y de primera clase (conceptos de la programación funcional), Go no es un lenguaje funcional, sino que se asemeja más a un lenguaje procedural como C.

Con Go se han creado programas como **Docker** y **Kubernetes**, los principales sistemas de contenedores para aplicaciones en la nube, y proyectos como **CockroachDB**, una base de datos SQL distribuida en la nube. También existen herramientas de alto nivel desarrolladas en este lenguaje, como **esbuild**, un empaquetador de proyectos de JavaScript extremadamente rápido, y **PocketBase**, un ejecutable que proporciona un backend como servicio similar a Firebase. Además, con **Wails** se pueden crear aplicaciones de escritorio utilizando Go para la lógica de backend y frameworks de frontend para las interfaces gráficas.

Go es una excelente opción para proyectos que requieren alta velocidad sin sacrificar la experiencia de desarrollo (DX).

Para hacer tu primer "Hola, mundo" en Go, por favor, ve a [este enlace](./src/README.md). Disfruta tu viaje en el maravilloso mundo de Go. ¡Mucha suerte, nuevo Gopher!
