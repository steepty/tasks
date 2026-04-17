// Вопросы: (теперь уже мои)
// 1) Что выведет принт?
// 2) В через разница defer и defer func? Как туда передаются значения? В какой момент

package main

import "fmt"

func main() {
	f := F(5)
	defer func() {
		fmt.Print(f())
	}()
	defer fmt.Print(f())
	f()
}

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}
