package main

import "fmt"



/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает:
три
 * 4) Напечатает:
два
 * 5) Напечатает:
один
 * 6) Напечатает:
пуск
 * 7) Напечатает:
один
пуск
*/

func main() {
	f([]int{}...)
}

func f(v ...int) {
	switch len(v) {
	case 3:
		fmt.Println("три")
	case 2:
		fmt.Println("два")
	case 1:
		fmt.Println("один")
		if v[0] < 0 {
			fallthrough
		}
	default:
		fmt.Println("пуск")
	}
}