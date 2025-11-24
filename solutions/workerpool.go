package solutions

import (
	"fmt"
	"sync"
)

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	for val := range jobs {
		results <- f(val)
	}
}

const numJobs = 5
const numWorkers = 3

func WorkerPoolMain() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}
	multiplier := func(x int) int {
		return x * 10
	}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(multiplier, jobs, results)
		}()
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		fmt.Println(val)
	}
}
