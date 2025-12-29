package main

import "fmt"

func main (){
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("Linha")

	pi := 3.14159265358979323846

	fmt.Printf("Usando print formatado: %.2f com duas casas decimais ou %.3f com 3 casas decimais", pi, pi)

	a:=1
	b := 1.99
	c := false
	d := "texto"
	
	fmt.Printf("\n%d %f %t %s", a, b, c, d)	
	fmt.Printf("\n %v, %v, %v, %v", a, b, c, d)
}
