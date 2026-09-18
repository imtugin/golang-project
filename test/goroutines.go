package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mutexs := &sync.Mutex{}
	var money, donationsCount int

	go func() {
		for {
			mutexs.Lock()
			if money != donationsCount {
				fmt.Println("money=", money, "donations=", donationsCount)
				break
			}
			mutexs.Unlock()
		}
	}()

	wg.Add(1000)
	for range 1000 {
		go func() {
			mutexs.Lock()
			defer wg.Done()
			money++
			donationsCount++
			mutexs.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(money)
}
