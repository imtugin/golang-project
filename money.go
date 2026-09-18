package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	_ "sync/atomic"
)

func main() {
	var money atomic.Int32
	wg := sync.WaitGroup{}
	wg.Add(3)
	for range 3 {
		go func() {
			defer wg.Done()
			money.Add(5)
		}()
	}
	wg.Wait()
	fmt.Println(money.Load())
}
