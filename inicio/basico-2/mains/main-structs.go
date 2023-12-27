package mains

//Retirar o arquivo que exite principal e renomear para main
import (
	"Golang/inicio/basico-2/model"
	"fmt"
)

func Structs() {

	//Exercicio items compra
	fmt.Println(model.IniciaCompraMes())

	/* 	endereco := model.Endereco{
	   		Rua:    "Rua dos Bobos",
	   		Numero: 0,
	   		Cidade: "Cidade do Bobo",
	   	}
	   	imovelCasa := model.Imovel{
	   		Tamanho:         "Grande",
	   		Valor:           1000000,
	   		NumeroDeQuartos: 2,
	   		EnderecoImovel:  endereco,
	   	}

	   	casa := model.Casa{
	   		Imovel: imovelCasa,
	   		Jardim: true,
	   	}

	   	apartamento := model.Apartamento{
	   		Imovel:         imovelCasa,
	   		Andar:          10,
	   		NumeroDeBlocos: 3,
	   		Elevador:       true,
	   	}

	   	fmt.Println(imovelCasa)
	   	fmt.Println(casa)
	   	fmt.Println(apartamento) */

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
