package solutions

import (
	"context"
	"fmt"
)

func GenSqrMain() {
	ctx := context.Background()

	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))

	for x := range pipeline {
		fmt.Println(x)
	}
}

func generator(ctx context.Context, in ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, value := range in {
			select {
			case <-ctx.Done():
				return
			case ch <- value:
			}
		}
	}()
	return ch
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	outChan := make(chan int)
	go func(){
		defer close(outChan)
		for  v := range in {
			select{
			case <-ctx.Done():
				return 
			case outChan <- v * v:
			}
		}
	}()

	return outChan
}
