package main

import (
	"fmt"
	"sync"
	_ "sync/atomic"
)

func main() {
	var money int32
	var donationsCount int32

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	go func() {
		for {
			mutex.Lock()
			m := money
			dc := donationsCount
			if m != dc {
				fmt.Println("money=", m, "donations=", dc)
				break
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
	wg.Wait()
	fmt.Println(money)

}
