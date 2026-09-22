package main

import (
	"fmt"
	"sync"
	_ "sync/atomic"
)

func main() {
	var money int
	var donationsCount int

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	wg.Add(1)
	go func() {
		for {
			mutex.Lock()
			m := money
			dc := donationsCount
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
			money++
			donationsCount++
		}()
		mutex.Unlock()
	}
	wg.Wait()
	fmt.Println(money)

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
