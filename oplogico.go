package main

import "fmt"

func main() {

	x := 5
	if x == 2 || x == 3 {
		fmt.Println("Sim, x é 2 ou 3!")
	} else {
		fmt.Println("X não é nem 2 e nem 3!")
	}
}
