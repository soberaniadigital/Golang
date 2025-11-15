/*
Código referente ao tutorial disponível no website Go By Example(https://gobyexample.com/switch)

# O código abaixo exemplifica o uso de arrays em Go

Autor: João Victor Oliveira
Data: 11-11-2025
*/
package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("array:", a)

	a[4] = 100
	fmt.Println("array:", a)
	fmt.Println("a[4] =", a[4])

	fmt.Println("tamanho: ", len(a))

}
