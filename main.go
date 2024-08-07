package main

import (
	t "banco-go/clientes"
	c "banco-go/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(float64) string
}

func main() {
	clienteLeandro := t.Titular{
		Nome:      "Leandro",
		CPF:       "12345678901",
		Profissao: "Desenvolvedor",
	}
	contaDeLeandro := c.ContaCorrente{Titular: clienteLeandro}
	contaDeGustavo := c.ContaCorrente{Titular: t.Titular{
		Nome:      "Gustavo",
		CPF:       "98765432109",
		Profissao: "Gerente",
	}}

	status := contaDeGustavo.Transferir(100, &contaDeLeandro)
	fmt.Println(status)
	fmt.Println(contaDeLeandro)
	fmt.Println(contaDeGustavo)
}
