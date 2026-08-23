package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Введите символ. Если введёте не 2, то должна появиться ошибка")
	fmt.Scanln(&a)
	switch a {
	case 2:
		fmt.Println("Вы ввели", a)
	default:
		fmt.Println("ошибка")
	}
}
