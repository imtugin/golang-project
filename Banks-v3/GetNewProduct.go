package main

import "fmt"

func RepeatGetNewProduct() {
	fmt.Println("Попробуйте ещё раз!")
}
func GetNewProduct(func(), error) {
	var choice int
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
			return GetCredit, nil
		case 2:
			return Save, nil
		case 3:
			return Inshurance, nil
		case 4:
			return GetDebitCard, nil
		case 5:
			return GetCreditCard, nil
		case 6:
			return Back, nil
		case 7:
			return Exit, nil
		}
	default:
		return RepeatGetNewProduct, parem_error{}

	}
}
