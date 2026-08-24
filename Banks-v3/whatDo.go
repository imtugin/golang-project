package main

import (
	"fmt"
	"os"
)

type parem_error struct{}

func (error_object parem_error) Error() string {
	return "Invalid parameter"
}

func RepeatWhatDo() {
	fmt.Println("Попробуйте ещё раз!")
}

func Exit() {
	fmt.Println("Выхожу из приложения")
	os.Exit(0)
}

func whatDo() (func() (func(), error), error) {
	var choice int
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
			return GetNewProduct, nil
		case 2:
			return MakeTransfer, nil
		case 3:
			return TopUpYourAccount, nil
		case 4:
			return View, nil
		case 5:
			return Exit, nil
		}

	default:
		goto metka
	}
}
