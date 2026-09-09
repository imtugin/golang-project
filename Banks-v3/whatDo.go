package main

import (
	"fmt"
	"os"
)

func RepeatWhatDo() {
	fmt.Println("Попробуйте ещё раз!")
}

func Exit() {
	fmt.Println("Выхожу из приложения")
	os.Exit(0)
}

func whatDo() {
	fmt.Println("Что хотите сделать?")
	fmt.Println("Получить новый продукт - 1")
	fmt.Println("Сделать перевод - 2")
	fmt.Println("Пополнить счёт - 3")
	fmt.Println("Посмотреть  'мои финансы' - 4")
	fmt.Println("Выйти - 5")
	fmt.Scanln(&choice)
metka:
	switch choice {
	case 1, 2, 3, 4, 5:

		switch choice {
		case 1:
			//			return GetNewProduct, nil
		case 2:
			MakeTransfer()
		case 3:
			TopUpYourAccount(&balance)
		case 4:
			//			return View, nil
		case 5:
			//			return Exit, nil
		}

	default:
		goto metka
	}
}
