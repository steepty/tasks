package main

import "fmt"

/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Ничего не напечатает и успешно завершится
 * 4) Напечатает:
привет как дела
 * 5) Напечатает:
<nil>
*/

func main() {
	if err := do(); err != nil {
		fmt.Println(err)
	}
}

func do() error {
	var p *customError
	return p
}

type customError struct{}

func (err customError) Error() string {
	return "привет как дела"
}
