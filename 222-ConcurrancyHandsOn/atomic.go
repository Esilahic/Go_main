package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)            //count is locked
			fmt.Println(atomic.LoadInt64(&count)) //count is unlocked and printed
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count:", count)
}
