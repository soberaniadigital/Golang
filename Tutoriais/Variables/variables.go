/*
	Código referente ao tutorial disponível no website Go By Example(https://gobyexample.com/variables)

	O código abaixo exemplifica o uso e inicialização de variáveis em Go

	Autor: João Victor Oliveira
	Data: 09-11-2025

*/
// Pacote main para definição de executável em go
package main

// O import "fmt" contém funções para formatação e saída de texto na E/S padrão
import "fmt"

// Função main
func main() {

	var a = "inicial" // declara a variável a e atribui o conteúdo "inicial" à variável
	fmt.Println(a)    // printa o conteúdo da variável passada como parâmetro.

	var b, c int = 1, 2 // declara as variáveis inteiras b e c e realiza inicialização
	fmt.Println(b, c)   // printa o conteúdo de b e c, respectivamente

	var d = true // declara variável booleana e inicializa com valor true(1)
	fmt.Println(d)

	var e int // declara e especificando o tipo mas sem inicializar. Nesse caso, o valor da varíavel inicializa com 0
	fmt.Println(e)

	// Abaixo, exemplo de declaração e inicialização curta de variável
	f := "maçâ"
	fmt.Println(f)
}
