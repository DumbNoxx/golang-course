# Primer hola mundo en golang

Generalmente esto es lo primero que haces cuando estas aprendiendo un lenguaje de programacion, hacer tu primer "Hello, World" en cualquier
lenguaje de programacion es un tradicion que se ha llevado por mas de 60 años, se podria decir que el primer "Hello, World" del mundo fue en C,
ya que es el mas conocido hasta la fecha, pero en realidad todo empezo con los lenguajes assembly (ensamblador) y mainframes en los años 1940 - 1950.

- **ENIAC (1945)**: Aunque no tenia una salida de texto como la que conocemos hoy, los programas se cableaban fisicamente para comunicarse.
- **IBM 701 (1952)**: Aunque se pasa de la fecha dicha anteriormente. Fue uno de los primeros sistemas donde programas en assembly imprimian caracteres en la pantalla de visualizacion basada en un tubo de rayos catodicos.

Ya por los años 1957 salio el primer lenguaje de alto nivel llamado **FORTRAN** o (Formula Translation) el cual esta diseñado para calculos cientificos.
Un ejemplo de una impresion en pantalla en FORTRAN ll (1958) podria verse asi:

```fortran
PROGRAM HELLO
  PRINT *, 'HELLO'
  END
```

No era un exactamente un "Hello, World", pero era el equivalente para la epoca, en diferencia a lo que tenemos hoy en dia.

Luego tenemos a COBOL (1959) que se **usa** para aplicaciones empresariales, mas enfocada en bancos, sistemas de procesamientos de transacciones, donde el manejo de
grandes volumenes de datos es crucial principalmente. Hice enfasis en el "usa" porque hasta el dia de hoy se sigue usando cobol en los bancos.

Luego tenemos a BCPL y B (1966 - 1969) que son los predecesores directos de C.

- BCPL: Es un lenguaje en el que se baso B. Un programa para imprimir texto era complejo para estandares modernos.
- B: Lenguaje creado por Ken Thompson en contribucion de Dennis Ritchie a finales de los años 60 fue una, fue desarrollado en los laboratorios Bell, donde mas tarde se crearia el sistema operativo UNIX 
Un ejemplo de una impresion en pantalla de B podria verse asi:
```b
main() {
    extrn a,b,c,d,e;
    putchar(a);putchar(b);putchar(c);putchar(d);putchar(e);
}
a 'hell';
b 'o, w';
c 'or';
d 'l';
e 'd'
```

Ya con esto se podria decir que llegamos directamente al famoso y mas conocido "Hello, World" el lenguaje C. Fue concebido como un lenguaje de programacion
que combinaba la eficiencia del assembly con la flexibilidad de un lenguaje de alto nivel. Su creador Dennis Ritchie tomo como inspiracion el lenguaje B, que a su vez
se basaba en BCPL. C se utilizo inicialmente para crear utilidades para el sistema operativo UNIX que originalmente fue escrito en assembly.
Una impresion en pantalla de C es de esta manera:

```c
main(){
  printf("Hello, World");
}
```

Con esto llegamos al primer "Hello, World" oficial, en ese vistazo de no usamos `#include<stdio.h>`. Por la
simple razon de que el pre-ANSI de c era mas permisivo, quiero decir que el compilador no requeria declaraciones explicitas de funciones como `printf`. Si la funcion
no se declaraba, el compilador asumia que devolvia un `int` y aceptaba cualquier numero de argumentos (el cual era un riesgo de bugs btw).
Ademas Kernighan (autor de "A Tutorial Introduction to the Language B" y co-autor de "The C Programming Language") y Dennis Ritchie querian enseñar la sintaxis basica sin
entrar en detalles de la librerias, ya en paginas posteriores del libro The C Programming Language explican que `printf` estaba definidada en `<stdio.h>`.
Llevando a que la sintaxis del lenguaje C llegaria a estar asi, hasta el dia de hoy:

```c
#include <stdio.h>

main() {
  printf("Hello, World");
}
```

Ya pasando esto, han habiado una infinidad de lenguajes de programacion en los cuales se ha hecho muy comun hacer un "Hello, World" en el lenguaje de programacion
que estas aprendiendo, porque es por asi decirlo una especie de introduccion a la sintaxis de aquel lenguaje.

Con esto en mente podemos pasar al primer "Hello, World" que haras en Go. [Click](./hello-world/README.md)
