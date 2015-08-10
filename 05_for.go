package main

import "fmt"

func main() {

	//un for simple
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//un for del tipo initial/condition/after
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	// un for sin condición se ejecutará hasta que
	// encuentre un break o un return
	for {
		fmt.Println("loop")
		break
	}
}
