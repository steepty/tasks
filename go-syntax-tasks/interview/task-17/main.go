// Вопросы: (теперь уже мои)
// 1) Что выведет каждый из принтов
// 2) Как под капотом, в виде структуры, выглядит интерфейс?
// 3) Если мы присваеваем пустому интерфейсу какую-то переменную, что внутри будет нил, а что нет? И в каких случаях будет будет разница

package main

import "fmt"

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil) 

	in = data
	fmt.Println(in, in == nil)

}
