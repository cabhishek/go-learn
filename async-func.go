package main

import (
	"fmt"
	"sync"
	"time"
)

func All(fns ...func()) (done <-chan bool) {
	var wg sync.WaitGroup

	wg.Add(len(fns))

	ch := make(chan bool, 1)

	for _, fn := range fns {
		go func(f func()) {
			f()
			wg.Done()
		}(fn)
	}

	go func() {
		wg.Wait()
		ch <- true
		close(ch)
	}()

	return ch
}

func PrintName() {
	time.Sleep(time.Second * 1)
	fmt.Println("Slim Shady")
}

func PrintGreetings() {
	fmt.Println("Hello world")
}

func main() {

	done := All(PrintName, PrintGreetings)

	<-done
}
