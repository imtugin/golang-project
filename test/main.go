package main

import (
	"fmt"
)

type param_error struct{}

func (error_object param_error) Error() string {
	return "Invalid parameter"
}

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, param_error{}
	}
	resoult := 1
	for i := 1; i < n; i++ {
		resoult *= 1
	}
	return resoult, nil
}

func main() {
	var f int
	fmt.Println("Введите число для факториала")
	fmt.Scanln(&f)
	f, err := factorial(f)
	fmt.Println("Factorial:", f)
	fmt.Println("Error", err)

}
