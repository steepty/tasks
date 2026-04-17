package main

import "fmt"

/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает:
1 2
1 2
0 0 0
 * 4) Напечатает:
1 2
10 20
0 1 0
 * 5) Напечатает:
1 2
1 2
0 1 0
 * 6) Напечатает:
1 2
1 2
[]
*/

func main() {
	fmt.Println(implicitly())
	fmt.Println(explicitly())
	fmt.Println(slicely())
}

func implicitly() (int, int) {
	var a, b int
	defer func() {
		a, b = 10, 20
	}()
	a, b = 1, 2
	return a, b
}

func explicitly() (a, b int) {
	defer func() {
		a, b = 10, 20
	}()
	a, b = 1, 2
	return
}

func slicely() []int {
	var out []int
	defer func() {
		out[1] = 1
	}()
	out = make([]int, 3)
	return out
}
