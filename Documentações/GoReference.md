By Evandro Policarpo
# Introduction
Este resumo, trata-se do manual de [referência](https://go.dev/ref/spec) da linguagem Go.
Go é uma linguagem de propósito geral, fortemente tipada, possui garbage collector e suporte explícito para programação concorrente. Também possui um sistema de packages para gerenciar dependências e sua sintaxe é simples e compacta, fácil para análise por ferramentar de CI/CD.

# Notation
A linguagem é especificada utiliazndo uma variante do `Extended Backus-Naur Form` (EBNF).
> Observação: a compreensão deste tópico está ligada à Linguagens Formais, Gramáticas e Teoria da Computação. O que não é muito prático, consequentemente também não é muito útil no momento.

# Source code representation
Códigos escritos em Golang são codificados em `UTF-8`. 
O texto não é canonicalizado, isto é: alguns caracteres, principalmente caracteres acentuados, podem ser representados de 2 formas diferentes: `Pré-composta` e `Decomposta`, por exemplo o ‘a’ acentuado (`á`), pode ser escrito como:
- Pré-composta: `á` (U+0061)
- Decomposta: `a + ´` (U+0061 + U0301)
Embora sejam 2 códigos diferentes, na prática eles representam o mesmo caractere, que é o a acentuado ‘`á`’ (Na maioria dos editores, a forma decomposta é “convertida” para a forma pré-composta, mas somente na exibição, no código fonte ele continua decomposto). 
O compilador de go enxerga diferença entre eles, a forma pré-composta é uma e a decomposta é outra diferente, ele não transforma as 2 em uma só.

- Obviamente, o compilador também percebe diferenças entre caracteres maiúsculos e minúsculos.
- O compilador proíbe o caractere NUL (U+0000) no código fonte.
- O compilador também pode proibir o `BOM` (byte order mark), exceto se estiver no início do arquivo.
  Byte order mark é um caractere unicode (U+FEFF) que indica que o texto está codificado em UTF-8, além de estar em Big-endian/Little-endian. Não é utilizado por editores modernos, e é utilizado por editores antigos somente no início do código.

>Resumindo: cada codepoint UTF-8 é distinto.

## Characters
Os caracteres são divididos por categorias:
- `newline` = code point U+000A.
- `unicode_char` = todos os caracteres unicode, com exceção do `newline`.
- `unicode_letter` = code point categorizado como “letra” pelo padrão unicode.
- `unicode_digit` = code point categorizado como “número” (decimal/dígito) pelo padrão unicode.

## Letters and digits
O caractere ‘\_’ é considerado como letra minúscula.
Sobre as 2 últimas categorias do tópico anterior, temos:
- `letra` = [unicode_letter](https://www.unicode.org/versions/Unicode8.0.0/) ou o caractere underscore `_`
- `decimal_digit` = ‘0’ … ‘9’
- `binary_digit` = ‘0’ ou ‘1’
- `octal_digit` = ‘0’ … ‘7’
- `hex_digit` = ‘0’ … ‘9’ ou ‘A’ … ‘F’ ou ‘a’ … ‘f’

# Lexical elements

## Comments
Servem para documentação/explicação do código.
Existem 2 tipos de comentários em Go (exatamente como em C):
- Single-line comment: `//comment...`
- Multi-line comment: `/* comment... */`
> Comentários não podem começar dentro de strings, runas ou de outros comentários. Um comentário sem quebras de linhas é tratado pelo compilador como um espaço em branco. Um comentário com pelo menos uma quebra de linha, é tratado pelo compilador como uma única quebra de linha (independente do comentário ter várias quebras de linha ou somente uma mesmo).

## Tokens
Um token é a maior sequência de caracteres que formam uma palavra válida em Go.
São basicamente as palavras que compõem um código em Go. Podem ser divididos em 4 categorias:
- identificadores
- palavras-chave
- operadores e pontuação
- literais
> Espaços em branco, tab, “retorno de carruagem” (\\r) e quebras de linha são ignorados, exceto quando eles separariam 2 tokens que formariam um único token

## Semicolons
No geral, linguagens de programação utilizam ponto e vírgula para encerrar uma linha, entretanto, em Go é possível omitir o ponto e vírgula seguindo 2 regras:
1. O compilador já insere automaticamente (por debaixo dos panos) o ponto em vírgula quando o último token da linha for um:
	- identificador
	- literal
	- palavra chave
	- operador ou pontuação.
2. Para permitir que instruções complexas ocupem uma única linha, o ponto e vírgula pode ser omitido antes de uma parêntese/chave fechada (`)` ou `}`)
> Basicamente, o compilador de Go já coloca o ponto e vírgula em todas declarações que fizermos automaticamente, **com exceção de múltiplas declarações por linha**, como em `i := 0 ; i++`,para executar as duas instruções (`i := 0` e `i++`) na mesma linha, é necessário o ponto e vírgula. 

## Identifiers
identificadores são utilizados para nomear entidades como variáveis e tipos. É uma sequência de um ou mais caracteres de letras ou dígitos, sendo que o primeiro caractere **deve** ser uma letra (lembrar do que foi escrito no primeiro tópico sobre caracteres).
- Exemplos de identificadores válidos: `a`, `_x9`, `helloworld`, `ぁ`.

## Keywords
São palavras/comandos já utilizados pelo compilador e por isso não podem ser utilizadas como identificadores:
- break
- case
- chan
- const
- continue
- default
- defer
- else
- fallthrough
- for
- func
- go
- goto
- if
- import
- interface
- map
- package
- range
- return
- select
- struct
- switch
- type
- var
> São exatamente 25 palavras-chave, o que torna Go uma das linguagens mais curtas do mundo.

## Operators and punctuation
São os operadores matemáticos, lógicos e de atribuição suportados pela linguagem.
- +
- -
- /
- *
- %
- &
- |
- ^
- << (shift-left)
- >> (shift-right)
- &^
- +=
- -=
- \*=
- \\=
- %=
- &=
- |=
- ^=
- <<=
- >>=
- &^=
- &&
- ||
- <-
- ++
- \--
- ==
- >
- <
- =
- !
- !=
- <=
- >=
- :=
- …
- ()
- {}
- \[]
- ,
- .
- ;
- :

## Integer literals
Como o nome sugere, representam valores inteiros. É prefixo opcional pode ser adicionado para representar bases numéricas não decimais, por exemplo:
- `0b` ou `0B` para binário
- `0o` ou `0O` para octal
- `0x` ou `0X` para hexadecimal
> É possível utilizar o `_` como separador.

Exemplos válidos de literais inteiros:
- 42
- 4_2
- 0600
- 0_600
- 0o600
- 0O600  
- 0xBadFace
- 0xBad_Face
- 0x_67_7a_2f_cc_40_c6
- 170141183460469231731687303715884105727
- 170_141183_460469_231731_687303_715884_105727

## Floating-Point literals
Pode ser representado utilizando decimal ou hexadecimal.
Decimal é composto por: 
1. parte inteira(‘0’ … ‘9’)
2. ponto decimal(`.`)
3. parte fracionária(‘0’ … ‘9’)
4. expoente(`e` ou `E`, seguido do sinal(opcional) `+` ou `-` e o expoente propriamente dito, na forma decimal)

Hexadecimal é composto por:
1. prefixo `0x` ou `0X`
2. parte inteira em hexadecimal (‘0’ … ‘F’)
3. ponto decimal(`.`)
4. parte fracionária(‘0’ … ‘F’)
5. expoente(`p` ou `P`, seguindo pelo sinal(opcional) `+` ou `-` e o expoente propriamente dito, na forma **decimal**).

> Exemplos:
- 1_5.  (`15.0`)
- 0.15e+0_2 (`15.0`)
- 0x1p-2 (`0.25`)
- 0x2.p10 (`2048.0`)

## Imaginary literals


## Rune literals
## String literals


# Constants

# Variables

# Types
## Boolean types
## Numerical types
## String types
## Array types
## Slice types
## Struct types
## Pointer types
## Function types
## Interface types
## Map types
## Channel types

# Properties of types and values
## Representation of values
## Underlying types
## Type identity
##  Assignability
## Representability
## Method sets


# Blocks

# Declarations and scope
## Label scopes
## Blank identifier
## Predeclared identifiers
## Exported identifiers
## Uniqueness of identifiers
## Constant declarations
## Iota
## Type declarations
## Type parameters declarations
## Variable declarations
## Short variable declarations
## Function declarations
## Method declarations


# Expressions
## Operands
## Qualified identifiers
## Compose literals
## Function literals
## Primary expressions
## Selectors
## Method expressions
### Method values
## Index expressions
## Slice expressions
## Type assertions
## Calls
## Passing arguments to parameters
## Instantiations
## Type inference
## Operators
## Arithmetic operators
## Comparison operators
## Logical operators
## Adress operators
## Receive operators
## Conversions
## Constant expressions
## Order of evaluation


# Statements
## Terminating statements
## Empty statements
## Labeled statements
## Send statements
## IncDec statements
## Assignment statements
## If statements
## Switch statements
## For statements
## Go statements 
## Select statements
## Return statements
## Break statements
## Continue statements
## Goto statements
## Falltrough statements
## Defer statements


# Bult-in functions
## Appending to and copying slices
## Clear
## Close
## Manipulating complex numbers
## Deletion of map elements
## Length and capacity
## Making slices, maps and channels
## Min and max
## Allocation
## Handling panics
## Bootstraping


# Packages
## Source file organization
## Package clause
## Import declarations
## An example package



# Program initialization and execution
## The zero value
## Package initialization
## Program initialization
## Program execution


# Errors


# Runtime panics

# System considerations
## Package unsafe
## Size and alignment guarantees

# Apendix
Language versions
Type unification rules
…


- [ ] Introduction
- [ ] Notation
- [ ] Source code representation
- [ ] Lexical elements
- [ ] Constants
- [ ] Variables
- [ ] Types
- [ ] Properties of types and values
- [ ] Blocks
- [ ] Declarations and scope
- [ ] Expressions
- [ ] Statements
- [ ] Bult-in functions
- [ ] Packages
- [ ] Program initialization and execution
- [ ] Errors
- [ ] Runtime panics
- [ ] System considerations
- [ ] Apendix
