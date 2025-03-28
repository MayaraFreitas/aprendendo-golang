package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	totalderoutines := 500
	mygoroutines(totalderoutines)
}

func mygoroutines(qty int) {
	wg.Add(qty)

	var contador int64
	contador = 0
	for i := 0; i < qty; i++ {

		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			wg.Done()
		}()
		//fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor Final:", contador)
}
