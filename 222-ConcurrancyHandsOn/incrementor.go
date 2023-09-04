package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0

	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() //count is locked
			v := count
			v++
			count = v
			fmt.Println(count)
			mu.Unlock() //count is unlocked
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count:", count)
}
