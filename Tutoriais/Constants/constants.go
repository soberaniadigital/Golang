/*
Código referente ao tutorial disponível no website Go By Example(https://gobyexample.com/constants)

# O código abaixo exemplifica o uso e inicialização de constantes em Go

Autor: João Victor Oliveira
Data: 09-11-2025
*/
package main

import (
	"fmt"
	"math"
)

const s string = "constante"

func main() {
	fmt.Println(s) // printa o conteúdo da constante s

	const n = 2000000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(d) // printa o conteúdo de d

	fmt.Println(int64(d)) // printa d como inteiro de 64 bits

	fmt.Println(math.Sin(d)) // faz chamado do método seno e printa o resultado

	fmt.Println(math.Sin(n))

}
