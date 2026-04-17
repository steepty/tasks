// Вопросы: (теперь уже мои)
// 1) Что выведет код?
// 2) Как работает чтение из мапы по ключу 

package main

import "fmt"

func main() {
	m := map[int]int{0: 0, 1: 10}
	fmt.Println(m, m[0], m[1], m[2])
}
