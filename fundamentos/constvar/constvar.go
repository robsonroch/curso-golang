package main

import (
	"fmt"
	"math"
)

func main(){
	const PI float64 = 3.1415
	var raio = 3.2 // tipo inferido pelo compilador

	// forma rduzida de criar uma variável
	area := 3.2 * math.Pow(raio, 2)

	println("Área da circunferência:", area)
	const (
		a =3
		b =4
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println("Constantes:", a, b, c, d)

	var e, f bool = false, true
	fmt.Println("variáveis",e, f)

	g, h, i := 1, false, "string"
	fmt.Println("variáveis 2:", g, h, i)
}
