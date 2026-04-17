package main


import (
	"fmt"
	"time"
)


/*
1
-
2
-
3
-
1
2
3
*/



/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает готово
 * 4) Ничего не напечатает и успешно завершится
 */

func main() {
  go func() {
    for i := 0; i < 10; i++ {
      time.Sleep(time.Millisecond*50)
    }
  }()
  select {}
  fmt.Println("готово")
}