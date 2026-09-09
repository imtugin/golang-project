package main

import (
	"fmt"
	"os"
)

func GetNewProduct() {
	var choice int
loop:
	fmt.Println("Что хотите сделать?")
	fmt.Println("Получить кредит - 1")
	fmt.Println("Накопить - 2")
	fmt.Println("Страхование - 3")
	fmt.Println("Получить дебетовую карту - 4")
	fmt.Println("Получить кредитную карту - 5")
	fmt.Println("Назад- 6")
	fmt.Println("Выйти - 7")
	fmt.Scanln(&choice)
	switch choice {
	case 1, 2, 3, 4, 5, 6, 7:
		switch choice {
		case 1:
			GetCredit()
		case 2:
			Save()
		case 3:
			Inshurance()
		case 4:
			GetDebitCard()
		case 5:
			GetCreditCard()
		case 6:
			whatDo()
		case 7:

			fmt.Println("Выхожу из программы...")
			os.Exit(3)
		}
	default:
		fmt.Println("Попробуйте ещё раз")
		goto loop
	}
}
