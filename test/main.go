package main

import (
	"fmt"
)

// создаём функцию факториал(перемножение всех чисел от 1 до n, n >1 должно быть, иначе ошибка)

func factorial(n int) (int, interface{}) {
	var resoult int = 1
	if n < 1 {
		return 0, "число должно быть не меншье 1"
	} else {
		for i := 1; i <= n; i++ {
			resoult *= i
		}
	}
	return resoult, nil
}
func main() {
	var f int
	fmt.Println("Введите число для факториала:")
	fmt.Scanln(&f)
	fmt.Println(factorial(f))

}
