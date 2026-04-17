package main

import "fmt"


/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Напечатает `я обязательно выживу`, а затем произойдёт ошибка runtime
 * 3) Напечатает:
я обязательно выживу
слишком сложно досвидания
 * 4) Программа не завершится
 */

const DATA_SIZE = 777 * 1024 // 777 KB

func main() {
  fmt.Println("я обязательно выживу")
  s := generate()
  recursion(s, len(s))
  fmt.Println("слишком сложно досвидания")
}

func generate() string {
  b := make([]byte, 0, DATA_SIZE)
  for i := 0; i < cap(b); i++ {
    b = append(b, byte(i))
  }
  return string(b)
}

func recursion(s string, i int) {
  if i == 0 {
    return
  }
  i--
  recursion(s, i)
}