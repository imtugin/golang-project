package main

import (
	"fmt"
	"sync"
)

func main() {
	var money, donationsCount int
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	go func() {
		for {
			defer wg.Done()
			mutex.Lock()
			m := money
			dc := donationsCount
			if m != dc {
				fmt.Println("money=", m, "donations=", donationsCount)
			}
			mutex.Unlock()
		}
	}()
	wg.Add(1000)
	for range 1000 {

		go func() {
			defer wg.Done()
			mutex.Lock()
			money++
			donationsCount++
			mutex.Unlock()
		}()
	}
	fmt.Println(money)
}
