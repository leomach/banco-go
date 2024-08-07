package contas

import (
	t "banco-go/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular                    t.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return fmt.Sprintf("Saque realizado com sucesso. saldo: %.2f", c.saldo)
	} else {
		return "saldo insuficiente para realizar o saque."
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso. saldo: %.2f", c.saldo
	} else {
		return "Valor do depósito não pode ser menor que zero.", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
