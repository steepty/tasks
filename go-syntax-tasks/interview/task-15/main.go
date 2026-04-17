package main

import "fmt"

/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает 13
 * 4) Напечатает 0
 * 5) Напечатает <nil>
 */

const TOTAL = 10_000

func main() {
	first := make(chan int)
	prev := first
	for i := 0; i < TOTAL; i++ {
		next := make(chan int)
		go func(prev <-chan int) {
			next <- <-prev
		}(prev)
		prev = next
	}
	first <- 13
	fmt.Println(<-prev)
}
