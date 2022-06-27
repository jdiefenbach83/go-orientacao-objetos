package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaCorrente1 := ContaCorrente{
		titular:       "Jefferson",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	contaCorrente2 := ContaCorrente{
		"José",
		555,
		111333,
		200.0,
	}

	fmt.Println(contaCorrente1)
	fmt.Println(contaCorrente2)
}
