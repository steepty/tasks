// Вопросы:
// 1) До запуска кода можно сразу сказать где ошибка? 
// 2) Почему так нельзя делать? Что может случиться?

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			fmt.Println("working")
		}()
	}

	wg.Wait()
}