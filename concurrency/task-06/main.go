// Вопросы:
// 1) Что здесь не так? Код конечно странный, но вроде же все работает. В чем проблема?
// 2) Это memory leak или нет? Почему?
// 3) Что делают горутины все это время?
// 4) Как такие проблемы ищут в коде?
package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	v := <-ch
	fmt.Println(v)
}

func main() {
	for i := 0; i < 100000; i++ {
		go worker(make(chan int))
	}

	time.Sleep(time.Second)
}