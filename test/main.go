package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var maxWaitSeconds = 5 // задаём максимальное время ожидания выполнения функции

func randomWait() int { // функция совершает работу какую-то
	//
	workSeconds := rand.Intn(5 + 1)                      //рандомная генерация времени выполнения которое быдет выполняться функция
	time.Sleep(time.Duration(workSeconds) * time.Second) //задерживаем функцию на это количество секунд
	return workSeconds

}

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	mutex := sync.Mutex{}
	var totalWorkSeconds int // время выполнения main (надо до 5 сек), сколько timesleep в сумме в каждом вызове
	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			seconds := randomWait()
			mutex.Lock()
			totalWorkSeconds += seconds
			mutex.Unlock()
		}()
	}
	wg.Wait()
	mainSeconds := time.Since(start)
	fmt.Println("main:", mainSeconds)
	fmt.Println("total:", totalWorkSeconds)
}
