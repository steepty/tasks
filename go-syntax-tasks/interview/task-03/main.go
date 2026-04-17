package main

import "fmt"

/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает 28 ¿
 * 4) Напечатает 15 р
 * 5) Напечатает 28 р
 */

func main() {
	greetings := "привет как дела"
	fmt.Printf("%d %c\n", len(greetings), greetings[1])
}
