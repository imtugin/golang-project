package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var money atomic.Int32
	var donationsCount atomic.Int32
	wg := sync.WaitGroup{}

	go func() {
		for {
			m := money.Load()
			dc := donationsCount.Load()

			if m != dc {
				fmt.Println("money=", m, "donations=", dc)
				break
			}
		}
	}()

	wg.Add(3)
	for range 3 {
		go func() {
			defer wg.Done()
			money.Add(5)
			donationsCount.Add(5)
		}()
	}
	wg.Wait()
	fmt.Println(money.Load())
	fmt.Println(donationsCount.Load())
}
