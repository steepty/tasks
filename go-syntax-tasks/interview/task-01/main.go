/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает 10 чисел от 1 до 10 в разном порядке
 * 4) Напечатает от 0 до 9 по порядку
 * 5) Ничего не напечатает и успешно завершится
 */
package main

import "fmt"

func main() {
  var ch chan int
  go func(c chan<- int) {
    for i := 0; i < 10; i++ {
      c <- i
    }
        close(c)
  }(ch)
  for v := range ch {
    fmt.Println(v)
  }
}
