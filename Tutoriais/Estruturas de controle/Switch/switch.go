/*
Código referente ao tutorial disponível no website Go By Example(https://gobyexample.com/switch)

# O código abaixo exemplifica a estrturua de controle switch

Autor: João Victor Oliveira
Data: 11-11-2025
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Escreva ", i, " como ")
	switch i {
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	case 3:
		fmt.Println("Três")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("É Fim de semana")
	default:
		fmt.Println("É um dia semana")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("É antes do meio-dia")
	default:
		fmt.Println("É depois do meio-dia")
	}

	quemSouEu := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Eu sou booleano")
		case int:
			fmt.Println("Eu sou inteiro")
		default:
			fmt.Printf("Não sei o tipo %T\n", t)
		}
	}

	quemSouEu(true)
	quemSouEu(1)
	quemSouEu("oi")

}
