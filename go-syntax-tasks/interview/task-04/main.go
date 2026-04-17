package main

import "fmt"

/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает [1 2 3 4 5]
 * 4) Напечатает [1 2 3 4 -1]
 * 5) Напечатает [1 2 3 4]
 */

const MAX = 5

func main() {
	s := generate()
	mutation(s)
	fmt.Println(s[0:MAX])
}

func generate() []int {
	out := make([]int, 0, MAX)
	for i := 1; i < MAX; i++ {
		out = append(out, i)
	}
	return out
}

func mutation(s []int) {
	s = append(s, -1)
}
