//Вопросы:
// 1) Что будет если один из воркеров зависнет? Поможет ли тут закрытие канала?
// 2) Нужен ли тут контекст, если да, то какой?
// 3) Как поправить код, чтобы не бояться зависания воркеров ?
// 4) Где должен проверяться `<-ctx.Done()` ?
package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int) {
	for job := range jobs {
		fmt.Println("processing", job)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	jobs := make(chan int)

	go worker(jobs)

	for i := 0; i < 5; i++ {
		jobs <- i
	}

	
}