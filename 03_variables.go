package main

import "fmt"

func main() {

	var a string = "initial"
	fmt.Println(a)

	//se pueden declarar mas de una variable a la vez
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//implicitamente una variable int se inicializa en 0
	var e int
	fmt.Println(e)

	//esta es la forma corta de una declaracion de un string
	f := "short"
	fmt.Println(f)
}
