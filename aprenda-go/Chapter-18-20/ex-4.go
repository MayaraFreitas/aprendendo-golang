package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	totalderoutines := 500
	mygoroutines(totalderoutines)
}

func mygoroutines(qty int) {
	wg.Add(qty)
	contador := 0
	var mu sync.Mutex

	for i := 0; i < qty; i++ {

		go func() {
			mu.Lock()
			v := contador
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		//fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor Final:", contador)
}
