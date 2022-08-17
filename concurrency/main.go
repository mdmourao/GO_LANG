package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func raceCondition() {
	var counter int64 = 0

	const gs = 100
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			// we could also do:
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
		fmt.Println("go routine test", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("COUNTER", counter)
}

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("go routine test", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("go routine test", runtime.NumGoroutine())
	wg.Wait()
	raceCondition()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
