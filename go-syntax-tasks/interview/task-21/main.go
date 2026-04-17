package main

import "fmt"

func main() {
	fmt.Println(add1(5)(3))
	fmt.Println(addAny(1, 1.5, 1.525))

	acc := adder()
	fmt.Println(acc(1))
	fmt.Println(acc(2))
	fmt.Println(acc(3))
}

func add1(a int) func(int) int {
	// 1) Написать реализацию, чтобы работал вызов add(5)(3) == 8
}

func addAny(nums ...any) any {
	// 2) Написать реализацию, чтобы была возможность одной функцией суммировать разные типы, например int и float
}

func adder() func(int) int {
	// 3) Классическая задача по замыканию ,если знаешь, будет проще.
	// Написать реализацию функции аккумулятора, которая сохраняет все предыдущие переданные значения,
}
