package contas

import (
	c "banco-go/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                              c.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return fmt.Sprintf("Saque realizado com sucesso. saldo: %.2f", c.saldo)
	} else {
		return "saldo insuficiente para realizar o saque."
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso. saldo: %.2f", c.saldo
	} else {
		return "Valor do depósito não pode ser menor que zero.", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}