package slicesarrays

import "fmt"

func maps() {
	/*nome_variavel := map[tipo_chave] tipo_valor
	m := map[string]int
	Em um 'map' não é possível adicionar novos elementos, mas é possível atualizar os valores existentes.
	e não podemos adicionar valores iguais

	chage valor [~list]
	*/

	map_variavel := map[string]string{"nome": "marcos", "idade": "20", "endereco": "endereco1"}
	for i, v := range map_variavel {
		fmt.Println(i, v)
	}

	/* adicionando items de um map  */
	map_variavel["sobrenome"] = "socram"

	for i, v := range map_variavel {
		fmt.Println(i, v)
	}

	/* deletando items de um map */
	delete(map_variavel, "nome")

}
