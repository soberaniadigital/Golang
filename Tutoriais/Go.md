https://gobyexample.com/
# Hello World
```Go
package main // Gera um executável
import "fmt"

func main(){ 
		fmt.Println("Hello World!")
}
```
Para compilar:
```shell
$ go run hello-world.go
  Hello World 
$ go build hello-world.go
$ ls
  hello-world 
  hello-world.go
$ ./hello-world
  "Hello World!"
```

`package` serve para modularizar o código, dividindo o código em pacotes de acordo com sua função. Se package for igual a `main`, significa que o programa é executável, caso contrário é somente biblioteca (não dá para executar nem com a função main).

# Values
**Inteiros:**
`int`, `int8`, `int16`, `int32`, `int64` `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `byte`, `rune`.
- `byte` é um alias para `uint8`.
- `rune` é um alias para `int32` (caractere unicode).
**Ponto flutuante:**
`float32` e `float64`
**Complexos:**
`complex64` e `complex128`
**Booleanos:**
`bool` (true ou false)
**Texto:**
`string`, `rune` e `byte` (lembrando: byte para caractere ASCII e rune para unicode)
- "Aspas duplas representa texto com caracteres de escape."
- 'Aspas simples representam taxtos literais (sem caracteres de escape)'
**Tipos compostos:**
- `array`, `slice`, `map`, `struct`, `pointer`, `function`, `interface`, `channel`.
**Tipos especiais:**
`nil`: vazio (o mesmo que NULL)
`structs`:
```Go
type Celsius float64
type Pessoa{
	Nome string
	Idad int8
}
```
Também é possível utilizar struct para fazer um alias de um tipo:
```Go
type MyInt = int
```

# Variables
- Variáveis podem ser declaradas com a palavra chave `var`. O compilador irá determinar automaticamente o tipo mais adequado.
- Também dá para criar uma variável com um tipo fixo, como em C, a diferença é que o tipo deve ser escrito depois do nome da variável, exemplo: 
  `var a int = 42`.
  `var b string = "Hello World"`
  `var c float16 = 3.141592`
- É possível criar múltiplas variáveis por declaração, mas também é necessário inserir múltiplos valores. Se quisermos criar as variáveis a,b e c com o valor 42, não dá para simplesmente colocar `var a,b,c = 42`, o correto é: `var a,b,c = 42,42,42`.
- Variáveis declaradas sem valor são inicializadas com valor padrão (0 para números, "" para String, false para bool...)
- Go possui uma sintaxe especial para criar e inicializar variáveis -> `:=` (mas não pode ser utilizado fora de funções):
  `f := 42`
```Go
package main

import "fmt"
func main() {
    var a = "initial"
    fmt.Println(a)
    var b, c int = 1, 2
    fmt.Println(b, c)
    var d = true
    fmt.Println(d)
    var e int
    fmt.Println(e)
    f := "apple"
    fmt.Println(f)
}
```
# Constants
- Define valores imutáveis (não podem ser alterados)
- Não possuem tipo definido. (É possível fazer casting para outros tipos tranquilamente)
- Expressões aritméticas com Const possuem precisão arbitrária.
- Não ocupa memória em tempo de execução.
```Go
package main
import (
    "fmt"
    "math"
)
const s string = "constant"
func main() {
    fmt.Println(s)
    const n = 500000000
    const d = 3e20 / n
    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))
}
```
```Shell
$ go run constant.go 
  constant
  6e+11
  600000000000
  -0.28470407323754404
```

# For
- É a única estrutura de loops em Go
- Funciona exatamente como em C: `for variável inicial ; condição de parada ; incremento {...}`
  Não possui parêntese
- Um for vazio é um loop infinito até que seja encontrada uma condição de parada DENTRO do loop.
- Go também suporta a sintaxe `range` de python: `for i := range 10 {...}`
```Go
package main
import "fmt"
func main() {
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }
    for i := range 3 {
        fmt.Println("range", i)
    }
    for {
        fmt.Println("loop")
        break
    }
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```

```Shell
$ go run for.go
  1
  2
  3
  0
  1
  2
  range 0
  range 1
  range 2
  loop
  1
  3
  5
```

# If/else
- Exatamente igual a C (Não é obrigatório colocar o parêntese).
- É possível realizar declarações dentro de cabeçalhos if-else. (Sendo que nesse caso, não é possível colocar parêntese)
- Não existe operador ternário em Go.
```Go
package main
import "fmt"
func main() {
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }
    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
```

```Shell
$ go run if-else.go
  7 is odd
  8 is divisible by 4
  either 8 or 7 are even
  9 has 1 digit
```

# Switch
- Exatamente igual a C
- É possível utilizar vírgula para separar múltiplas expressões na mesma declaração.
```Go
package main
import (
    "fmt"
    "time"
)
func main() {
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
```

```Shell
$ go run switch.go 
  Write 2 as two
  It's a weekday
  It's after noon
  I'm a bool
  I'm an int
  Don't know type string
```

# Arrays
- É uma sequência fixa de elementos do mesmo tipo.
- Diferente de C, arrays em Go são inicializados automaticamente com valores padrão.
- É possível criar um array (necessariamente já inicializado) utlizando \[...], que irá contar a quantidade de elementos inclusos no array.
- Basta colocar a quantidade de elementos do array entre colchetes (caso ele não seja inicializado manualmente) **antes** do tipo de elemento.
```Go
package main
import "fmt"
func main() {
    var a [5]int
    fmt.Println("emp:", a)
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    fmt.Println("len:", len(a);
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
    var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}
```

```Shell
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
dcl: [1 2 3 4 5]
idx: [100 0 0 400 500]
2d:  [[0 1 2] [1 2 3]]
2d:  [[1 2 3] [1 2 3]]
```

# Slices
- Fazem parte do pacote "slices"
- É um conjunto dinâmico de elementos do mesmo tipo.
- É basicamente uma lista.
- É possível criar um slice a partir da função `make` (basicamente make retornaria uma struct lista, se fosse em C):
```Go
var n = make([]int, length, capacity);
//O tipo pode ser []string, []float... qualquer tipo, mas TEM que ter o colchete.
```
- Fatiamento: é possível fatiar um trecho de um slice, criando assim, um slice menor (Mas que ainda aponta para o mesmo array, se ele for deletado, o slice que foi fateado também será deletado).
```Go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]  // slice dos índices 1,2,3
fmt.Println(s) // [2 3 4]
```
- [Funcionamento interno(por debaixo dos panos) de um slice](https://go.dev/blog/slices-intro)

```Go
package main

import (
    "fmt"
    "slices"
)

func main() {

    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

    twoD := make([][]int, 3)
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```

```Shell
$ go run slices.go
  uninit: [] true true
  emp: [  ] len: 3 cap: 3
  set: [a b c]
  get: c
  len: 3
  apd: [a b c d e f]
  cpy: [a b c d e f]
  sl1: [c d e]
  sl2: [a b c d e]
  sl3: [c d e f]
  dcl: [g h i]
  t == t2
  2d:  [[0] [1 2] [2 3 4]]
```

# Maps

by Evandro <3
