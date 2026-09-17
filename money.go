package main

import (
	"fmt"
	"sync"
)

func main() {
	var money int
	wg := sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			money++
		}()
		wg.Wait()
	}
	fmt.Println(money)
}
