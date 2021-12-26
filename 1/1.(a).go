package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		mu.Lock()

		fmt.Println("有点强人锁男")

		mu.Unlock()
		wg.Done()
	}()
	wg.Wait()

	//源代码的mu.Unlock在主协程中，应该在子协程Unlock之后才能进行，所以Unlock应该放在子协程中

}
