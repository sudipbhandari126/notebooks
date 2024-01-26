package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total uint64

func worker(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total, i)
	}
	ch <- int(total)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int, 2)

	go worker(&wg, ch)
	go worker(&wg, ch)
	wg.Wait()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
