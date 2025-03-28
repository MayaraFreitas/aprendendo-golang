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
	for i := 0; i < qty; i++ {

		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		//fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor Final:", contador)
}
