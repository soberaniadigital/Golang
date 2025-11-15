/*
Código referente ao tutorial disponível no website Go By Example(https://gobyexample.com/for)

# O código abaixo exemplifica o uso do laço de repetição for

Autor: João Victor Oliveira
Data: 09-11-2025
*/
package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i) // printa de 1 a 3
		i = i + 1
	}
	// printa de 0 a 2
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	// Printa de 0 a 2
	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	// Printa os 6 primeiros números pares
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
