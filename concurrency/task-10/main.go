// Вопросы:
// 1) Почему тут дедлок?
// 2) Как поправить чтобы дедлок не было?
// 3) Нужен ли `defer` в горутине?
package main

import (
	"fmt"
)

func main() {
	sem := make(chan struct{}, 3)

	for i := 0; i < 10; i++ {
		sem <- struct{}{}

		go func(i int) {
			fmt.Println("job", i)
		}(i)
	}
}
