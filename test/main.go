package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var money atomic.Int32
	var donationsCount atomic.Int32

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	wg.Add(1)
	go func() {
		for {
			mutex.Lock()
			m := money.Load()
			dc := donationsCount.Load()
			fmt.Println(m, dc)
			if m != dc {
				fmt.Println("money=", m, "donations=", dc)
				break
			}
			mutex.Unlock()
		}
	}()
	wg.Add(1000)
	for range 1000 {
		mutex.Lock()
		go func() {
			defer wg.Done()
			money.Add(1)
			donationsCount.Add(1)
		}()
		mutex.Unlock()
	}
	wg.Wait()
	fmt.Println(money.Load())

}

//func main() {
//	var money int
//	wg := &sync.WaitGroup{}
//	wg.Add(1000)
//	for range 1000 {
//		go func() {
//			defer wg.Done()
//			money++
//		}()
//	}
//	wg.Wait()
//	fmt.Println(money)
//
//}
