package conc

import (
	"fmt"
	"time"
)

func MainTask3() {
	ch := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Отдельнная горутаина запущена")
		ch <- false
	}()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Произошел новый тик")
			ch <- true

		case value := <-ch:
			fmt.Println("Пришло значние", value)
			return
		}
	}
}
