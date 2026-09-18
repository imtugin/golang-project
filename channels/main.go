package main

import (
	"fmt"
)

func foo(n int, c chan int) {
	n += 10
	c <- n
}
func main() {
	ch := make(chan int)
	var slice []int
	var b int
	a := 5

	go func() {
		a += 10
		ch <- a
		a *= a
		ch <- a
	}()
	a, b = <-ch, <-ch
	fmt.Println(a, b)

	go func() {
		a += 13
		ch <- a
		ch <- 30
		ch <- 60
		ch <- 6000
	}()
	read, a, b, z := <-ch, <-ch, <-ch, <-ch
	fmt.Println(read, a, b, z)
	go func() {
		for range 6 {
			a *= a
			ch <- a
		}
	}()
	for i := 0; i <= 6; i++ {
		slice[i] = <-ch
	}
	fmt.Println(slice)

}
