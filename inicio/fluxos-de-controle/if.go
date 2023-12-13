package fluxosdecontrole

import "fmt"

func main(){

	salario := 1001.00
	salarioMaior := 1300
	var salarioMaisBonus float64

	salarioMaisBonus = salario

	if salario < 1000.0 && salarioMaior <= 1200 {
		salarioMaisBonus = salario + 100
	} else if salario > 1000.0 && salarioMaior >= 1200 {
		salarioMaisBonus = salario + 300
	}

	fmt.Println(salarioMaisBonus)
}
