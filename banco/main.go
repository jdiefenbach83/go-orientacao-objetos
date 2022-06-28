package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaCorrente1 := ContaCorrente{}

	contaCorrente1.titular = "Jefferson"
	contaCorrente1.saldo = 500

	fmt.Println(contaCorrente1.saldo)

	fmt.Println(contaCorrente1.Sacar(300))
	fmt.Println(contaCorrente1.saldo)
	fmt.Println(contaCorrente1.Sacar(600))
	fmt.Println(contaCorrente1.Sacar(-100))
}
