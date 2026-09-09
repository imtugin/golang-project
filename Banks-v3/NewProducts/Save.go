package main

import "fmt"

func Save() {
loop:
	fmt.Println("Выберите из предложенного:")
	fmt.Println("Накопительный счёт - 1")
	fmt.Println("Депозит - 2")
	fmt.Println("Цифровой актив - 3")
	fmt.Println("Назад - 4")
	fmt.Println("Выйти - 5")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		Deposit()
	case 2:
		SavingAccount()
	case 3:
		DigitalAssets()
	case 4:
		GetNewProduct()
	default:
		fmt.Println("Попробуйте ещё раз")
		goto loop
	}
}
