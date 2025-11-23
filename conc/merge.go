package conc

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func merge(ctx context.Context, channels ...chan int64) chan int64 {
	res := make(chan int64)

	merger := func(ch chan int64) {
		for{
			select{
			case <-ctx.Done():
				fmt.Println("Завершение контекста")
				return 
			case val, ok := <-ch:
				if !ok{
					return 
				}
				res <-val
			}
		}
	}

	wg := sync.WaitGroup{}

	wg.Add(len(channels))
	for _, ch := range channels {
		go func(){
			defer wg.Done()
			merger(ch)
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()
	return res
}

func MergeMain() {
	channels := make([]chan int64, 10)
	for i := range channels {
		channels[i] = make(chan int64)
	}

	for i := range channels {
		go func(i int) {
			channels[i] <- int64(i)
			close(channels[i])
		}(i)
	}
	ctx , cancel:= context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for v := range merge(ctx,channels...) {
		println(v)
	}

	ch := make(chan int64)
	merge(ctx, ch ,ch)
}