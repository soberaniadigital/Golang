/*
Código referente ao tutorial disponível no website Go By Example(https://gobyexample.com/switch)

# O código abaixo exemplifica a estrutura Slice

Autor: João Victor Oliveira
Data: 12-11-2025
*/
package main

import (
	"fmt"
)

func main() {

	var s []string
	fmt.Println("unidade: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len: ", len(s), "capacidade:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("conjunto: ", s) // imprime todos os componentes
	fmt.Println("s[2]", s[2])    // imprime o terceiro elemento

	fmt.Println("len:", len(s))

	s = append(s, "d") // faz chamada do método append para adicionar elementos
	s = append(s, "e", "f")
	fmt.Println("Conjunto: ", s)

	c := make([]string, len(s))
	copy(c, s) // copia slice s

	l := s[2:5]            // copia elementos dos indices de 2 a 5 de s
	fmt.Println("sl1:", l) // imprime l

}
