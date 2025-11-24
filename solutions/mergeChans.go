package solutions

import (
	"fmt"
	"sync"
)

func fillChan(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()
	return ch
}

func merge(cs ...<-chan int) <-chan int {
	mergedCh := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	for _, tempCh := range cs {
		go func(tempCh <-chan int) {
			defer wg.Done()
			for v := range tempCh {
				mergedCh <- v
			}
		}(tempCh)
	}

	go func() {
		wg.Wait()
		close(mergedCh)
	}()
	return mergedCh
}

func Task1Main() {
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)

	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}
}
