package main

import "fmt"

func main() {
	x := 5
	y := &x // & é o operador de endereço de memória
	*y = 10 // * é o operador de desreferência
	fmt.Println("Na main")
	fmt.Println(x, *y)
	fmt.Println(&x, &y)
	//imprimirValores(x, * y)

}

func imprimirValores(x int, y int) {

}
