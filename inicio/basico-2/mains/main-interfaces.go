package mains

import (
	"Golang/model"
	"Golang/repositorio"
	"fmt"
)

func Interfaces() {
	/* 	//Interfaces

	   	circulo := model.IniciaRetangulo()
	   	retangulo := model.IniciaCirculo()

	   	ExibeGeometria(circulo)
	   	ExibeGeometria(retangulo) */

	enderecos := model.InicializaEnderecos()

	enderecosNoRepositorio := repositorio.EnderecoRepositorio{
		Enderecos: enderecos,
	}

	id := 3
	enderecoEncontrado := enderecosNoRepositorio.FindByID(id)
	fmt.Println(enderecoEncontrado, ": enderecoEncontrado")

	nomeRua := "Travessa das Flores"
	cidadeEncontrada := enderecosNoRepositorio.FindByRua(nomeRua)
	fmt.Println(cidadeEncontrada, ": enderecoEncontrado 2 ")

}

/*
func ExibeGeometria(g interfaces.Geometria) {
	fmt.Println(g.Area())
}
*/
