package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func ParallelSum(s []int) int {
	cores := runtime.NumCPU()
	partSize := len(s) / cores

	c := make(chan int, cores)

	for i := 0; i < cores; i++ {
		start := i * partSize
		end := start + partSize
		if i == cores-1 {
			end = len(s) // последний кусок включает остаток
		}
		go sum(s[start:end], c)
	}

	total := 0
	for i := 0; i < cores; i++ {
		total += <-c
	}
	return total
}

func JustSum(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	// Генерируем данные
	slice := make([]int, 100_000_000)
	for i := range slice {
		slice[i] = rand.Int()
	}

	fmt.Println("Ядер:", runtime.NumCPU())

	// Обычное суммирование
	start := time.Now()
	result1 := JustSum(slice)
	fmt.Println("Без горутин:", time.Since(start), "Результат:", result1)

	// Параллельное суммирование
	start = time.Now()
	result2 := ParallelSum(slice)
	fmt.Println("С горутинами:", time.Since(start), "Результат:", result2)
}
