package main

import (
	"fmt"
	"os"
)

// создать функцию сделать её результатом другой функции. и разбить по модулям.
// в общем будет выбор что сделать - сложить, вычесть или умножить два числа

func add(x, y int) int       { return x + y }
func substract(x, y int) int { return x - y }
func multiply(x, y int) int  { return x * y }

func selectFn(n int) (func(int, int) int, error) {
	switch n {
	case 1:
		return add, nil
	case 2:
		return substract, nil
	case 3:
		return multiply, nil
	default:
		return , nil
	}
}
