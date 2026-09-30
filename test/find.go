package main

import (
	"fmt"
)

func main() {
	var n int
	var lenSl int
	fmt.Println("Срез от 1 до скольки?")
	fmt.Scanln(&lenSl)
	fmt.Println("Максимальное количество попыток при бинарном поиске может достигать", MaxAttempts(lenSl))
	fmt.Println()
	sl := MakeSl(lenSl)
	fmt.Println("Введите число для поиска от 1 до 240000")
	fmt.Scanln(&n)
	fmt.Println(find(n, sl))
}

func MakeSl(l int) []int {
	var slice []int
	for i := 0; i < l; i++ {
		slice = append(slice, i+1)
	}
	return slice
}

// эта функция просто ищет число и выводит количество попыток за которое оно найдено
// но возможно будет ошибка если количество элементов в срезе будет не чётным, хотя может int это решит
// И есть проблема - поиск всегда проходит за 18 попыток. Я что не могу найти число раньше?
func find(x int, a []int) (bool, int) {
	var isFind bool
	var nubmersOfAttempts int

	for i := 1; ; i++ {
		halfSlice := len(a) / 2
		fmt.Println()
		fmt.Println("Попытка:", i)
		fmt.Println("Поиск из", halfSlice, "Чисел")
		if x == a[halfSlice] {
			isFind = true
			nubmersOfAttempts = i
			break
		} else if x > a[halfSlice] {
			a = a[halfSlice:]
			fmt.Println("между", a[0], "и", a[len(a)-1])
		} else {
			a = a[:halfSlice]
			fmt.Println("между", a[0], "и", a[len(a)-1])
		}
	}
	return isFind, nubmersOfAttempts
}

// но если длина массива будет 0, то будет бесконечно проводиться поиск
func MaxAttempts(l int) int {
	var maximum int
	if l == 0 {
		fmt.Println("Здесь не где искать")
		maximum = 0
	}
	if maximum != 0 {
		for i := 1; ; i++ {
			comparsion := l
			l /= 2
			fmt.Println("------------------------------")
			fmt.Printf("половина равна:\t\t%d\n", l)
			if l*2 != comparsion {
				l += 1
				fmt.Printf("прибавляем единицу:\t%d\n", l)
			}
			if l == 1 {
				maximum = i
				break
			}
		}
	}

	fmt.Println("==========================================================================")
	return maximum
}
