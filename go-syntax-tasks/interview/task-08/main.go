package main

import "fmt"




/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает:
uno, dos, tres, cuatro, cinco, cinco, seís
 * 4) Напечатает:
uno, dos, tres, cinco, cinco, seís
 * 5) Напечатает:
cuatro, cinco, cinco, seís
 * 6) Напечатает:
uno, dos, cuatro, cinco, cinco, seís
 * 7) Напечатает:
cinco, cinco, seís
*/

func main() {
LOOP:
	for i, s := range []string{"uno, ", "dos, ", "tres, "} {
		if i > 2 {
			goto LABEL
		}
		if i < 3 {
			continue LOOP
		}
		fmt.Print(s)
	}
	fmt.Print("cuatro, ")
LABEL:
	fmt.Println("cinco, cinco, seís")
}