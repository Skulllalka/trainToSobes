package mymap

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache sync.Map

func GetOrCompute(key string, computeFunc func() string) string {
	if val, ok := cache.Load(key); ok {
		fmt.Printf("Взяли из кэша %s\n", key)
		return val.(string)
	}
	newValue := computeFunc()

	cache.Store(key, newValue)

	return newValue
}

func compute(userID string) string {
	time.Sleep(1 * time.Second)
	fmt.Printf("Сложные вычисления для %s\n", userID)

	return fmt.Sprintf("result %s", userID)
}

func Task5() {
	usersID := []string{"user1", "user2", "user1", "user2", "user3", "user4", "user7"}
	wg := sync.WaitGroup{}

	for _, userID := range usersID {
		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(10))* time.Second)
			res := GetOrCompute(userID, func() string {
				return compute(userID)
			})
			fmt.Println(res)
		}()
	}

	wg.Wait()
}
