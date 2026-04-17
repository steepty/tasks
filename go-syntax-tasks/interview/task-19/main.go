// Вопросы: (теперь уже мои)
// 1) Что выведет принт?
// 2) Что за двоеточия в скобке? Что они дают?
// 3) Что дает `::`? Почему последний принт не работает? Как починить?

package main

import "fmt"

func main() {
	numbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := numbers[:5:8]
	fmt.Println(s1)
	s2 := numbers[:7:8]
	fmt.Println(s2)
	fmt.Println(numbers[::4]) 
}
