// Вопросы:
// 1) Почему здесь дедлок?
// 2) Что поправить в коде, чтобы дедлока не было?
// 3) Кто здесь должен закрыть канал?
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
