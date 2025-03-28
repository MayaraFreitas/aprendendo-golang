package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	minhagoroutine(2)
	wg.Wait()

}

func minhagoroutine(qty int) {

	wg.Add(qty)

	for i := 0; i < qty; i++ {
		x := qty
		go func(qty int) {
			fmt.Println("goroutine-", i+1)
			wg.Done()
		}(x)
	}

}
