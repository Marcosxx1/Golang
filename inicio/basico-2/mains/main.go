package mains

import "fmt"

func Generics() {
	var intSlice = []int{10, 11, 12, 13, 14, 15, 16}
	var stringSlice = []string{"a", "b", "c", "d"}
	reverse(intSlice)
	reverse(stringSlice)
}

type CustomType interface { // constraints
	int | string | float32
}

func reverse[T CustomType](slice []T) { // função genérica
	newInts := make([]T, len(slice))

	aux := 0
	for i := len(slice) - 1; i >= 0; i-- {
		newInts[aux] = slice[i]
		aux++
	}
	fmt.Println(newInts)
}
