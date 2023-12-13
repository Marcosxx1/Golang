package fluxosdecontrole

import "fmt"

func ExemploFor() {

	for  i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	
	/* Percorre uma matriz 2x3  */
	var a [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = i + j
			if(a[i][j] == 2) {
				break
			}
			fmt.Print(a[i][j], " ")
			continue
		}
	}

	text := "palavra"
	for i, v := range text {
		fmt.Printf("%d %c\n", i, v)
	}

	/* Percorre um map */
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	for k, v := range m {
		fmt.Printf("%s %d\n", k, v)
	}

	/* Percorre um slice */
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range s {
		fmt.Printf("%d %d\n", i, v)
	}

	/* Percorre um string */
	str := "palavra"
	for i, v := range str {
		fmt.Printf("%d %c\n", i, v)
	}

	/* Percorre um array */
	textoPalavra := "palavra"
	for i := 0; i < len(textoPalavra); i++ {
		fmt.Println(string(textoPalavra[i]))
	}

	/* Percorre um array usando range*/
	textoPalavra = "palavra"
	for _, v := range textoPalavra {
		fmt.Println(string(v))
	}

	/*
	for range é um for com índice e valor que é mais apropriado 
	para percorrer slices, arrays, strings e maps.

	for i, v := range textoPalavra {
		fmt.Println(i, v)

		i é a posição (índice)
		v é o valor (caractere)

	} 
	
	podemos também ignorar o índice:
	for _, v := range textoPalavra {
		fmt.Println(v)
	}

	podemos também ignorar o valor :
	for i := range textoPalavra {
		fmt.Println(i)
	}
	*/
}