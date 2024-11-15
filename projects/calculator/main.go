package main

import (
	"fmt"
	"sync"
)

func main() {
	firstch := make(chan int)
	secondch := make(chan int)
	stopch := make(chan struct{})
	ch := calculator(firstch, secondch, stopch)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for ans := range ch {
			fmt.Println(ans)
		}
		wg.Done()
	}()
	secondch <- 11
	wg.Wait()
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	res := make(chan int)
	go func(res chan int) {
		defer close(res)
		select {
		case x := <-firstChan:
			res <- x * x
		case x := <-secondChan:
			res <- x * 3
		case <-stopChan:
		}
	}(res)
	return res
}
