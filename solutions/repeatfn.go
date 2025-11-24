package solutions

import (
	"context"
	"math/rand"
)

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- fn():
			}
		}
	}()
	return ch
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	outCh := make(chan interface{})

	go func(n int) {
		defer close(outCh)

		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok{
					return 
				}
				outCh<- value
			}

		}
	}(num)

	return outCh
}

func RepeatFnMain() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rand := func() interface{} {
		return rand.Int()
	}

	var res []interface{}

	for num := range take(ctx, repeatFn(ctx, rand), 3) {
		res = append(res, num)
	}

	if len(res) != 3 {
		panic("wrong code")
	}
}
