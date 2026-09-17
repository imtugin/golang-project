package main

import (
	"fmt"
	_ "sync"
	"time"
)

func main() {
	var timer int
	for range 1000 {
		timer += 1
		fmt.Println(timer)
		time.Sleep(1000000000)
	}

	fmt.Println("Таймер завершился")
}
