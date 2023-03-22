package main

import "fmt"

const ebulicaoK = 373.0

func main() {
	tempK := ebulicaoK
	tempC := tempK - 273.0

	fmt.Printf("A temperatura da ebulição da agua em K = %g , temperatura da agua em °C = %g.", tempK, tempC)
}
