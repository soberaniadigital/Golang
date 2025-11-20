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
- Estrutura de dados map (ou dicionário/dict...) importada no pacote `"map"`.
- Sabemos que um `map` é um par de valores {chave, valor}. Para criar um map, utilizamos a mesma função `make` com o tipo da chave entre colchetes, seguido pelo tipo do valor, por exemplo:
```Go
m := make(map[string]int) // Cria um map do tipo {string:int}
n := make(map[int]int)    // Cria um map do tipo {int:int}
x := map[string]int{"Hello World" : 42} // Cria e inicializa um map.
```
- Para acessar um valor, é semelhante a acessar um elemento de um array atraveś da variável do array e o index do elemento,basta trocar o index pela chave. Exemplo:
```Go
m := make(map[string]int){"Hello" : 7, "World" : 42}
m["Hello"]; // Está acessando 7
m["World"]; // Está acessando 42
m["INEXISTENTE"]; // Está acessando 0.
```
>Observação: caso tentemos acessar uma chave que não foi inicializada (ou não existe), obteremos o resultado padrão de variáveis não inicializada do valor daquele map. No exemplo acima, como tratava-se de um int, retornou 0.
- Temos algumas funções auxiliares para estruturas de dados, nos quais se destacam `delete(map, key)` que recebe um map no primeiro argumento e uma chave desse map, no segundo argumento e deleta a chave e seu respectivo valor. `clear(map)` é como dar um free no map (se aplica para todos os seus elementos, automaticamente).
- Operador blank `_` é utilizado para ignorar um valor, ver `Exemplo 1` no código abaixo->
 
```Go
package main
import (
    "fmt"
    "maps"
)
func main() {
    m := make(map[string]int)
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m)
    v1 := m["k1"]
    fmt.Println("v1:", v1)
    v3 := m["k3"]
    fmt.Println("v3:", v3)
    fmt.Println("len:", len(m))
    delete(m, "k2")
    fmt.Println("map:", m)
    clear(m)
    fmt.Println("map:", m)
    _, prs := m["k2"] //Exemplo 1
    fmt.Println("prs:", prs)
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}
```

```Shell
$ go run maps.go 
  map: map[k1:7 k2:13]
  v1: 7
  v3: 0
  len: 2
  map: map[k1:7]
  map: map[]
  prs: false
  map: map[bar:2 foo:1]
  n == n2
```

> Observação: Fazendo uma analogia com Java é como se estas estruturas de dados (slice, map, array) herdassem de uma superclasse "Data Structure", por exemplo, que implementa os métodos length, print, delete e algumas outras funções.
> Em outras palavras, cada uma destas estruturas compartilha um mesmo conjunto de métodos auxiliares, porém cada uma possui sua própria implementação do método, por exemplo: o método print funciona diferente para um array (que imprime todos os elementos) e para um map (que imprime os pares chave/valor), embora continuem sendo uma função print.

>Observação 2: se for criar um map sem inicializar ele automaticamente, utilize **SEMPRE** a função make. (Caso contrário o map será null e tentativas de acesso resultarão em um "null pointer exception")

# Functions
- Funções são criadas com a palavra chave `func` seguida pelo nome da função, seus parâmetros e, por fim, seu tipo de retorno.
- Os parâmetros são escritos entre parêntese e separados por vírgulas (como em C). Diferente de C, o tipo do parâmetro deve ser escrito depois do nome. Se todos os parâmetros forem do mesmo tipo, é possível colocar um único tipo após todos eles (para não ter que ficar repetindo).
```Go
package main
import "fmt"
func plus(a int, b int) int {
    return a + b
}
func plusPlus(a, b, c int) int {
    return a + b + c
}
func main() {
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
```

```Shell
$ go run functions.go 
  1+2 = 3
  1+2+3 = 6
```

# Multiple Return Values
- É possível retornar múltiplos valores em Go, basta colocar os respectivos tipos dos retornos entre parêntese (no final do cabeçalho da função e antes das chaves)
- Para recolher este retorno múltiplo, basta separar as variáveis que irão receber o valor do retorno por vírgula e realizar uma atribuição com `:=`, como no exemplo abaixo:
```Go
package main
import "fmt"
func vals() (int, int) {
    return 3, 7
}
func main() {
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
    _, c := vals()
    fmt.Println(c)
}
```

```Shell
$ go run multiple-return-values.go
  3
  7
  7
```

# Variadic Functions (Var args Function)
- É uma função com quantidade variável de argumentos(ou parâmetros).
- Basta incluir 3 pontos `...` antes do tipo do parâmetro. Isto irá tornar a quantidade dele variável, na prática ele irá ser tratado como um array (com todas as propriedades do array, como len, inclusive)
```Go
package main
import "fmt"
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
func main() {
    sum(1, 2)
    sum(1, 2, 3)
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

```Shell
$ go run variadic-functions.go 
  [1 2] 3
  [1 2 3] 6
  [1 2 3 4] 10
```

# Closures
- É uma função que retorna uma função. A função maior (aquela que contém uma função no seu escopo) armazena as variáveis da função menor.
- Utiliza passagem por referência.
- É difícil de entender, vejamos os exemplos:
```Go
package main
import "fmt"
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}
```

```Shell
$ go run closures.go
  1
  2
  3
  1
```

```Go
package main
import "fmt"
func main() {
    counter := func() func() int {
        x := 0
        return func() int {
            x++
            return x
        }
    }()
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
```

# Recursão
- Possui a mesma lógica que C. Funções anônimas (closures) também podem ser recursivas, sendo que exigem que a variável seja declarada explícitamente com var para armazenar a função antes dela ser definida.
```Go
package main
import "fmt"
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
func main() {
    fmt.Println(fact(7))
    var fib func(n int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    fmt.Println(fib(7))
}
```

```Shell
$ go run recursion.go 
  5040
  13
```

# Range over Built-in Types
- `range` itera sobre todos os elementos de estruturas de dados embutidas (array, slice, map, strings...).
- Em estruturas com index como `arrays` e `slices` range tem 2 retornos: o índice e o elemento do índice.
- Em `maps` ele também retorna 2 valores: `chave` e `valor`.
- É como se fosse uma função de 2 retornos que recebe uma estrutura de dados embutida e percorre todos os seus elementos, imprimindo o index/valor ou chave/valor.
- Como `range` retorna 2 valores, aqui temos um bom exemplo de como aplicar o `blank` identifier (\_), caso nós tenhamos interesse somente em um dos valores do retorno, colocamos o valor que não nos interessa como `blank`.
```Go
package main

import "fmt"
func main() {
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    for k := range kvs {
        fmt.Println("key:", k)
    }
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```

```Shell
$ go run range-over-built-in-types.go
  sum: 9
  index: 1
  a -> apple
  b -> banana
  key: a
  key: b
  0 103
  1 111
```

# Pointers
É extremamente parecido com C, na verdade é idêntico no que tange a: 
- Um ponteiro de qualquer tipo, é uma variável que armazena o endereço de memória de alguma variável (do mesmo tipo do ponteiro)
- Símbolos: \* é utilizado para indicar um ponteiro e & para endereços.
- É possível criar variáveis do tipo ponteiro, utilizando u \* antes do tipo, ex: `var n *int`
Porém é diferente nos seguintes aspectos:
- Go possui um garbage collector, então o `malloc()` e `free()` acontecem de forma automática.
- Não é possível fazer casting de ponteiros.
- Não existe aritmética de ponteiros.

```Go
package main
import "fmt"
func zeroval(ival int) {
    ival = 0
}
func zeroptr(iptr *int) {
    *iptr = 0
}
func main() {
    i := 1
    fmt.Println("initial:", i)
    zeroval(i)
    fmt.Println("zeroval:", i)
    zeroptr(&i)
    fmt.Println("zeroptr:", i)
    fmt.Println("pointer:", &i)
}
```

```Shell
$ go run pointers.go
  initial: 1
  zeroval: 1
  zeroptr: 0
  pointer: 0x42131100
```

# Strings & Runes
https://go.dev/blog/strings
- Diferente das outras linguagens, em que uma String é um array de caracteres (inteiro de 1 byte), Strings em go são um array de runas.
- Runas são caracteres UNICODE(UTF-8). Acontece que os caracteres menores como os ASCII são representados com apenas 1 byte e os outros caracteres (acentuados, emojis, outros alfabetos) são representados com mais bytes (algo entre 2 e 4 bytes...).
- Por isso que são representados como um array de `runas` e não um array de `caracteres utf-8`, já que cada runa pode ter entre 1-4 bytes e se utilizássemos `caracteres utf-8` teriamos que ter EXATAMENTE 8 bytes (o que seria um desperdício de espaço).
- Em Go Strings não possuem um caractere '\0' para representar o fim da string.
- Strings em Go são imutáveis, não dá para mudar um caractere dps que a string é criada, é preciso recriar a string completamente.
- Para contar o tamanho de uma String utilizamos 2 funções: `len` que retorna a quantidade de bytes da String (cada runa pode ter mais de um byte, então len != quantidade de caracteres) e `range` (que conta, de fato, a quantidade de caracteres).
```Go
package main
import (
    "fmt"
    "unicode/utf8"
)
func main() {
    const s = "สวัสดี"
    fmt.Println("Len:", len(s))
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()
    fmt.Println("Rune count:", utf8.RuneCountInString(s))
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
        examineRune(runeValue)
    }
}
func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}
```

```Shell
$ go run strings-and-runes.go
  Len: 18
  e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5 
  Rune count: 6
  U+0E2A 'ส' starts at 0
  U+0E27 'ว' starts at 3
  U+0E31 'ั' starts at 6
  U+0E2A 'ส' starts at 9
  U+0E14 'ด' starts at 12
  U+0E35 'ี' starts at 15

  Using DecodeRuneInString
  U+0E2A 'ส' starts at 0
  found so sua
  U+0E27 'ว' starts at 3
  U+0E31 'ั' starts at 6
  U+0E2A 'ส' starts at 9
  found so sua
  U+0E14 'ด' starts at 12
  U+0E35 'ี' starts at 15
```

> Lembrando que `rune` é um alias para `int32` e `byte` para `uint8`.

# Structs
- São declaradas com a palavra chave `type` seguida do nome, `struct` e as chaves.
- São iguais a C no sentido de armazenar variáveis de diversos tipos.
- Os campos da Struct podem ser acessados utilizando o nome da struct, seguido pelo operador `.` seguido do nome do campo.
- Não existe operador `->` em Go.

```Go
package main
import "fmt"
type person struct {
    name string
    age  int
}
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}
func main() {
    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name: "Alice", age: 30})
    fmt.Println(person{name: "Fred"})
    fmt.Println(&person{name: "Ann", age: 40})
    fmt.Println(newPerson("Jon"))
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
    sp := &s
    fmt.Println(sp.age)
    sp.age = 51
    fmt.Println(sp.age)
    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}
```

```Shell
$ go run structs.go
  {Bob 20}
  {Alice 30}
  {Fred 0}
  &{Ann 40}
  &{Jon 42}
  Sean
  50
  51
  {Rex true}
```

# Methods
- Um método é uma função que pertence a uma classe. Go não tem classes, então métodos são funções que pertencem à structs.
- Na prática, a função será declarada por fora, mas pertencerá à struct.
- Sintaticamente, basta adicionar (entre parêntese) uma variável seguida pelo nome da estruct antes do nome da função. É como se estivesse adicionando a struct como um parâmetro da função. Exemplo:
```Go
package main
import "fmt"
type rect struct {
    width, height int
}
func (r *rect) area() int {
    return r.width * r.height
}
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
func main() {
    r := rect{width: 10, height: 5}
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
```

```Shell
$ go run methods.go 
  area:  50
  perim: 30
  area:  50
  perim: 30
```

# Interfaces
- Uma interface é um conjunto de funções que deve ser aplicado para todas as classes filhas. Como Go não tem classes, então aplicamos interfaces em structs.
- A sintaxe é a mesma de structs, só muda o nome de `struct` para `interface` e ao invés de ter atributos, tem-se funções e seus respectivos tipos.
- Para uma struct "implementar" uma interface, basta implementar todos os métodos dela.
```Go
package main
import (
    "fmt"
    "math"
)
type geometry interface {
    area() float64
    perim() float64
}
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius)
    }
}
func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    measure(r)
    measure(c)
    detectCircle(r)
    detectCircle(c)
}
```

```Shell
$ go run interfaces.go
  {3 4}
  12
  14
  {5}
  78.53981633974483
  31.41592653589793
  circle with radius 5
```

# Enums
- Não são embutidos por padrão na linguage, mas é possível implementá-los artificialmente utilizando o identificador especial `iota`.
- `iota` é uma variável constante que começa em 0 e é incrementada para cada variável subsequente do bloco const. No exemplo abaixo `Debug` é inicializado com 0, `Info` (A variável consecutiva) consequentemente será 1, `Warning` 2 e `Error` 3.
### Exemplo Simples:
```Go
package main

import "fmt"

type LogLevel int

const (
	Debug LogLevel = iota //0
	Info                  //1
	Warning               //2
	Error                 //3
)

func main() {
	fmt.Println(Debug, Info, Warning, Error)
}
```
Saída:
```0 1 2 3```
### Exemplo Complicado:
```Go
package main
import "fmt"
type ServerState int
const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)
var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}
func (ss ServerState) String() string {
    return stateName[ss]
}
func main() {
    ns := transition(StateIdle)
    fmt.Println(ns)

    ns2 := transition(ns)
    fmt.Println(ns2)
}
func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:

        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
```
Saída:
```Shell
$ go run enums.go
  connected
  idle
```

# Struct Embedding
- É possível criar structs dentro de outras structs, criando assim structs embarcadas. (Exatamente como em C, caso a struct seja criada com typedef).
- Diferente de C, é possível adicionar uma struct já existente como um campo de uma nova struct, portanto, todas as modificações feitas na struct menor, afetarão a struct maior, já que não será criada uma cópia da struct menor. É como se passassemos um ponteiro de struct como campo.
```Go
package main
import "fmt"
type base struct {
    num int
}
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}
type container struct {
    base
    str string
}
func main() {
    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
    fmt.Println("also num:", co.base.num)
    fmt.Println("describe:", co.describe())
    type describer interface {
        describe() string
    }
    var d describer = co
    fmt.Println("describer:", d.describe())
}
```
Saída:
```Shell
$ go run struct-embedding.go
  co={num: 1, str: some name}
  also num: 1
  describe: base with num=1
  describer: base with num=1
```
# Generics
- Permitem escrever funções, métodos e tipos que funcionam com vários tipos de dados, sem precisar fazer uma implementação para cada tipo (como é em C, geralmente).
- Sintaxe: `func Sum[T int | float64](a, b T) T {...}`
	- `func` descreve que é uma função (como já vimos)
	- `Sum` é o nome da função (como também já vimos)
	- `[T int | float64]` está dizendo que o tipo `T` (que acabamos de criar) pode ser `int` ou `float64`. (Se quiséssemos que T fosse de QUALQUER tipo utilizaríamos `any`: `[T any]`)
	- `(a, b T)` declara os parâmetros da função, que são `a` e `b`, ambos do tipo `T` (recém declarado, que pode ter `int` ou `float64`).
	- Este último `T` implica que o retorno da função é do tipo `T` (que, mais uma vez, pode ser `int` ou `float64`).

Exemplo:
```Go
package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}

type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    var s = []string{"foo", "bar", "zoo"}

    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

    _ = SlicesIndex[[]string, string](s, "zoo")

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}
```

```Shell
$ go run generics.go
  index of zoo: 2
  list: [ 10 13 23 ]
```

>Observação: embora todos os exemplos usem a letra `T` para representar o tipo, ela não é uma palavra-chave, é apenas uma convenção e pode ser substituida por qualquer outra letra/palavra.


# Range over Iterators
Explicação melhor sobre `range`:
- Percorre estruturas que são iteráveis, no caso: arrays, slice, string e maps.
- Sintaxe: `for index, value := range []int{10,20,30} {...}`
	- index é a posição atual.
	- value é o valor daquela posição.
	- range percorrerá os 3 elementos daquele array, um elemento para cada iteração.
	Observação: `value` é uma cópia do elemento do array, elementos muito grandes, vou ocupar um espaço extra muito grande (overhead).

- range retorna os seguintes valores de `value` de acordo com o tipo de estrutura que ele está percorrendo:
	- Slice/Array: retorna o elemento na posição `index`, simples assim.
	- Strings: retorna a runa na posição `index`. (Não confundir com o byte!!)
	- Maps: retorna um par `chave` e `valor` em vez de um único `value`.

Sobre `iterators`:
- Permitem que tipos customizados possam ser percorridos com range e não somente os fornecidos pela linguagem.
- Para utilizar iteradores em tipos compostos, basta criar uma função que retorne dois valores: 1 de qualquer tipo para representar o dado que está sendo acessado e 1 booleano para definir se continua a iteração ou para.
Exemplo:
```Go
func SliceIter[T any](s []T) func() (T, bool) {
    i := 0
    return func() (T, bool) {
        if i >= len(s) {
            var zero T
            return zero, false
        }
        v := s[i]
        i++
        return v, true
    }
}
```

Exemplo um pouco mais complexo:
```Go
package main

import (
    "fmt"
    "iter"
    "slices"
)

type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) All() iter.Seq[T] {
    return func(yield func(T) bool) {

        for e := lst.head; e != nil; e = e.next {
            if !yield(e.val) {
                return
            }
        }
    }
}

func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 1, 1

        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}

func main() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)

    for e := range lst.All() {
        fmt.Println(e)
    }

    all := slices.Collect(lst.All())
    fmt.Println("all:", all)

    for n := range genFib() {

        if n >= 10 {
            break
        }
        fmt.Println(n)
    }
}
```
Resultado:
```Shell
$ go run range-over-iterators.go
  10
  13
  23
  all: [ 10 13 23 ]  
  1
  1
  2
  3
  5
  8
```

# Errors
- Pertencem ao pacote “errors”.
- Como é possível criar funções com múltiplos retornos em Go, geralmente um destes retornos é do tipo erro, caso venha a ocorrer, senão é retornado NULL (ou NIL) na posição do erro. Na prática erros acabam se comportando exatamente como tipos de dados.
- É possível criar novos erros (como variáveis) utilizando a função `New` do pacote `errors`, passando apenas a mensagem de erro como parâmetro da função.
- O pacote `fmt` possui uma função `Errorf` que recebe uma String e trata o formatador `%w` para `errors.`
- O pacote `errors` também possui uma função `Is` que recebe 2 erros e retorna true se eles forem iguais. (Ela compara 2 erros, que nem `strcmp` da stdlib).

```Go
package main

import (
    "errors"
    "fmt"
)
func f(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")
func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {
        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}
func main() {
    for _, i := range []int{7, 42} {
        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }
    for i := range 5 {
        if err := makeTea(i); err != nil {
            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }
        fmt.Println("Tea is ready!")
    }
}
```

``` Shell
$ go run errors.go
  f worked: 10
  f failed: can't work with 42
  Tea is ready!
  Tea is ready!
  We should buy new tea!
  Tea is ready!
  Now it is dark.
```

# Custom Errors
- Erros customizados em Go, significam um novo tipo (criado com a palavra-chave type) que implementa a função Error(). Vejamos:
```Go
type MyCustomError struct{
	Resorce string;
}

func (x MyCustomError) Error() string{
	return fmt.Sprintf("%s not found", x.Resource);
}

//Usage:
func FindUser(id int) error{
	return MyCustomError{Resource : "User"}
}

var err error = FindUser(10);
if err != nil{
	fmt.println(err); // ""User" not found".
}
```
- Observações: por padrão utilizamos ponteiros para o tipo para evitar cópias desnecessárias. Ou seja: deveria ser `(x *MyCustomError)`.

Exemplo mais complexo:
```Go
package main
import(
	"errors"
	"fmt"
)

type argError struct{
	arg int
	msg string
}

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.msg);
}

func f(arg int) (int, error){
	if (arg == 42){
		return -1, &argError(arg, "can't work with it");
	}
	return arg + 3, nil;
}

func main(){
	_, err := f(42);
	var ae *argError;
	if(error.As(err, &ae)){
		fmt.Println(ae.arg);
		fmt.Println(ae.msg);
	}else{
		fmt.Println("err doesn't match argError");
	}
}
```
Saída:
```Shell
$ go run custom-errors.go
  42
  can't work with it
```

# Goroutines
- É uma maneira simples e eficiente de criar threads em Go. Não são exatamente threads ao pé da letra, mas goroutines permitem que funções diferentes executem concorrentemente (ao mesmo tempo).
- Basta utilizar `go` antes da declaração da função que criará uma goroutine para a função.
- Diferente de threads, goroutines são extremamente leves, elas começam com 2kb e vão aumentando conforme a necessidade. (Uma thread de SO começa com, no mínimo, 1MB). Além disso threads são gerenciadas pelo SO enquanto `goroutines` são gerenciadas pelo runtime de Go.
```Go
package main
import(
	"fmt"
	"time"
)
func f(str string){
	for i:= range 3{
		fmt.Println(string,":",i);
	}
}
func main(){
	f("direct"); //Irá executar f totalmente e só depois irá seguir para o resto do código.
	
	go f("goroutine"); //Irá executar uma iteração do loop e depois seguirá para o resto do código.
	
	go func(msg string){
		fmt.Println(msg);
	}("Going");
	
	time.Sleep(time.Second);
	fmt.Println("Done");
}
```

```Shell
$ go run goroutines.go
  direct : 0
  direct : 1
  direct : 2
  goroutine : 0
  going
  goroutine : 1
  goroutine : 2
  done
```

# Channels
- São os “pipes" (ou tubulações) que conectam 2 `goroutines` diferentes. É o que permite o compartilhamento e dados entre 2 goroutines diferentes
- Cria-se uma channel com `make` e `chan`, de acordo com o tipo de dado que o channel transporta.
Exemplo:
```Go
msg := make(chan string);
msg2 := make(chan int);
//...
```
- Utiliza-se a sintaxe `<-` tanto para enviar quanto para acessar dados do chanel. 
```Go
package main
import(
	"fmt"
)
func main(){
	messages := make(chan string);
	
	go func(){
		messages <- "ping"; //Atribui ao channel messages a string "ping".
	}() // chama a função imediatamente após criar.
	
	msg := <-messages //Cria uma variável msg e atribui-lhe o conteúdo de messages, ou seja: "ping".
	fmt.Println(msg); //"ping".
}
```

# Channel Buffering

# Channel Synchronization

# Channel Directions
