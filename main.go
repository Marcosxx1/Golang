package main

import (
	"fmt"
	"os"
)

func MostraTexto() {
	fmt.Println("Finalizando")
}

func main() {

	// Cria o arquivo
	file, err := os.Create("./setting.go")

	// Verifica por erros
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte("teste"))
	if err != nil {
		panic(err)
	}

	// Finaliza a E/S de arquivos
	defer file.Close() // defer =
	// Depois mostra o texto
	defer MostraTexto()
}
