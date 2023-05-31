package main

import "fmt"

func main() {
	var a int
	var b float64 = 373.2

	a = int(b) - 273

	fmt.Println("Ponto de ebulição da água em Kelvin: ", b,
		"\nPonto de ebulição da água em Celsius: ", a)
}
