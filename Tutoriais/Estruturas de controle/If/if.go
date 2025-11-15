/*
Código referente ao tutorial disponível no website Go By Example(https://gobyexample.com/if)

# O código abaixo exemplifica o uso da estrutura de controle switch

Autor: João Victor Oliveira
Data: 11-11-2025
*/
package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 ou 7 são pares")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "é negativo")
	} else if num < 10 {
		fmt.Println(num, "tem um único dígito")
	} else {
		fmt.Println(num, "tem múltiplos dígitos")
	}

}
