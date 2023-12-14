package main

import "Golang/model"

func main() {

	imovelCasa := model.Casa {
    Tamanho: "Grande",
    Valor: 1000000,
    NumeroDeQuartos: 2,
    Endereco: model.Endereco{
      Rua: "Rua dos Bobos",
      Numero: 0,
      Cidade: "Cidade do Bobo",
    },
  }


  }

	/*   arquivoJson := model.Endereco{
	    Rua:    "Rua Meu Endereço",
	    Numero: 21,
	    Cidade: "Cidade pequena",
	  }

	  pessoa := model.Pessoa{
	    Nome:     "João",
	    Idade:    0,
			DataNascimento: time.Date(2000, 11, 25, 0,0,0,0, time.Local),
	    Endereco: arquivoJson,
	  }


	  fmt.Println("pessoa: ", pessoa)
		pessoa.CalculaIdade()
	  fmt.Println("idade: ", pessoa.Idade)
	  arquivoJson.Apresentar() */

}