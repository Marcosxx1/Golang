package mains

import (
	"fmt"
	"os"
)

func ReadFile() {

		defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado")
		}
	}() // função anonima

	_, err := os.ReadFile("./settings.txt")
	if err != nil {
		panic("FILE-NOT-FOUND") // para a execução
	}

}

func PanicDefer() {

	ReadFile()

	fmt.Println("FIM")

}

/* 	// Cria o arquivo
file, err := os.Create("./setting.go")

// Verifica por erros
if err != nil {
	panic("FILE-DOES-NOT-EXIST")
}

_, err = file.Write([]byte("teste"))
if err != nil {
	panic(err)
}

// Finaliza a E/S de arquivos
defer file.Close() // defer =
// Depois mostra o texto
defer MostraTexto() */
