package main

import (
	"fmt"
	"os"
)

func (object_error object_error) Error() string {
	return "no many!"
}

type param_error interface {
	Error() string
}
type object_error struct{}

func Transfer(b *float64, s float64) interface{} {
	if s <= *b {
		*b -= s
		return nil
	} else {
		return object_error{}
	}
}

func MakeTransfer() {
	var sumTransfer float64
	fmt.Println("Введите сумму перевода")
	fmt.Scanln(&sumTransfer)
	err := Transfer(&balance, sumTransfer)
	if err != nil {
		fmt.Println("Не достаточно средств")
	} else {
		fmt.Println("Перевод осуществлён")
		fmt.Println("На балансе:", balance)
	}
	fmt.Println("Вернуться на главный экран - y. Новый перевод - n. Выйти - exit")
	fmt.Scanln(&strChoice)
	switch strChoice {
	case "y":
		whatDo()
	case "n":
		MakeTransfer()
	case "exit":
		fmt.Println("Выхожу из программы...")
		os.Exit(0)
	default:
		fmt.Println("Выбран неизвестный вариант. Возвращаюсь на главный экран...")
		whatDo()
	}

}
