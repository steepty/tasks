//Вопросы:
// 1) Чем отличаются Inc1 и Inc2?
// 2)На что влияет `*`? 
// 3) В чем разница pointer ресивера и value ресивера? Когда нужен один, а когда другой?
package main

import "fmt"

type Counter struct {
	n int
}

func (c Counter) Inc1() {
	c.n++
}

func (c *Counter) Inc2() {
	c.n++
}
func main() {
	var c Counter
	c.Inc1()
	fmt.Println(c.n)
	c.Inc2()
	fmt.Println(c.n)
}
