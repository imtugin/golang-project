package main

import "fmt"

func Transfer() {
	fmt.Println("Перевод осуществлён как будто")
}
func MakeTransfer() (func(), error) {
	fmt.Println("Повторите ещё раз!")
	fmt.Println("Кому перевести")
	return Transfer, nil
}
