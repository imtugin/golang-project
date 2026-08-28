package main

import "fmt"

func TopUpYourAccount(b *float64) {
	var sumTopUp float64
	fmt.Println("Введите сумму, на которую хотите пополнить")
	fmt.Scanln(&sumTopUp)
	*b += sumTopUp
	fmt.Println("Пополнение осуществлено")
	fmt.Println("На балансе", balance)
	whatDo()
}
