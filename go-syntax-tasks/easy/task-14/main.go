// Вопросы:
// 1) Что вообще за `defer`? Почему порядок выполнения именно такой?
// 2) Что общего у `defer` и стека?
package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
