// fibonacci.go
// Название: Fibonacci Generator
// Краткое описание: Модуль для генерации последовательности чисел Фибоначчи.
// Разработчик: Ваше Имя
// Лицензия: GPLv3
// История изменений: 1.0 - первая версия

package main

import (
  "fmt"
  "math/rand"
  "time"
)

// GenerateFibonacci генерирует последовательность Фибоначчи и отправляет ее в канал
func GenerateFibonacci(ch chan<- int, n int) {
  a, b := 0, 1
  for i := 0; i < n; i++ {
    ch <- a
    a, b = b, a+b
  }
  close(ch)
}

// PrintFibonacci получает числа из канала и выводит их на экран с добавлением случайного числа
func PrintFibonacci(ch <-chan int) {
  for num := range ch {
    randomNum := rand.Intn(10) + 1 // Генерируем случайное число от 1 до 10
    fmt.Printf("%d + %d = %d\n", num, randomNum, num+randomNum)
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())
  ch := make(chan int)

  go GenerateFibonacci(ch, 10) // Генерируем 10 чисел Фибоначчи
  PrintFibonacci(ch)            // Печатаем числа
}
