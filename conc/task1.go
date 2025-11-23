package conc

import (
	"context"
	"log"
	"math/rand"
	"sync/atomic"

	"time"
)

var counter atomic.Int64

func SimulateRequest(ctx context.Context) (int64, error) {
	start := time.Now()
	defer func() {
		log.Printf("Время выполнения запроса: %v\n", time.Since(start))
	}()

	ch := make(chan int64)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		counter.Add(1)
		ch<- counter.Load()
	}()

	select{
	case <-ctx.Done():
		return 0, ctx.Err()
	case count:= <- ch:
		return count, nil
	}
	
	

}

func PseudoMain() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	value , err:= SimulateRequest(ctx)
	if err!= nil{
		log.Printf("Ошибка выполнения запроса: %v\n", err)
		return
	}
	log.Printf("Значени счётчика: %d\n", value)
}
