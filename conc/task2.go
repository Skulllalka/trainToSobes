package conc

import (
	"context"

	"fmt"
	"time"
)

const defaultTimeout = 1 * time.Second

func getDiscount() float64 {
	time.Sleep(2 * time.Second)
	return 12.0
}

func getDiscountWithTimeout(ctx context.Context) (float64, error) {
	ch := make(chan float64)
	go func() {
		res := getDiscount()
		ch <- res
		close(ch)
	}()

	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
	}

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case value := <-ch:
		return value, nil
	}

}

func Task2Main() {
	ctx := context.Background()

	res, err := getDiscountWithTimeout(ctx)
	if err != nil {
		fmt.Printf("Ошибка %v", err)
		return
	}
	fmt.Println("Ваша скидка: ", res)
}
