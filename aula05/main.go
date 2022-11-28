package main

import "fmt"

func main() {

}

var produtoI interface {
	preco() float32
}

type produtoPequeno struct {
	preco float32
}

type produtoMedio struct {
	preco float32
}

type produtoGrande struct {
	preco float32
}

func (produto produtoPequeno) calcularPrecoProdutoPequeno() float32 {
	return produto.preco
}

func (produto produtoMedio) calcularPrecoProdutoMedio() float32 {
	return produto.preco * 1.03
}

func (produto produtoGrande) calcularPrecoProdutoGrande() float32 {
	return produto.preco * 1.06 + 50
}

func newProduto(tipoProduto string, preco float32) produtoI {
	return nil
}