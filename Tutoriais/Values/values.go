/*
	Código referente ao tutorial disponível no website Go By Example(https://gobyexample.com/values)

	O código abaixo exemplifica tipos de dados e suas manipulações em Go

	Autor: João Victor Oliveira
	Data: 09-11-2025

*/

// Pacote main para definição de executável em go
package main

// O import "fmt" contém funções para formatação e saída de texto na E/S padrão
import "fmt"

// Função main
func main() {

	fmt.Println("go" + "lang")          // concatenação de strings
	fmt.Println("1 + 1 =", 1+1)         // soma de tipos inteiro
	fmt.Println("7.0 / 3.0 =", 7.0/3.0) // manipulação de tipos ponto flutuante

	// Tipos booleanos/lógicos
	fmt.Println(true && false) // resulta em false
	fmt.Println(true || false) // resulta em true
	fmt.Println(!true)         // resulta em false

}
