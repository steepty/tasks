//Вопросы:
// 1) Почему теперь изменения не повлияли на `a`?
// 2) В каких случаях нужно использовать copy() , а в каких передавать через присваивание (b := a)?

package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)

	b[0] = 100

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}