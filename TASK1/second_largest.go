package main

import (
  "fmt"
  "math"
)

// FindSecondLargest находит второй по величине элемент в массиве
func FindSecondLargest(arr []int) (int, error) {
  if len(arr) < 2 {
    return 0, fmt.Errorf("массив должен содержать как минимум два элемента")
  }

  max := math.MinInt
  secondMax := math.MinInt

  for _, num := range arr {
    if num > max {
      secondMax = max
      max = num
    } else if num > secondMax && num < max {
      secondMax = num
    }
  }

  if secondMax == math.MinInt {
    return 0, fmt.Errorf("второй по величине элемент не найден")
  }

  return secondMax, nil
}

func main() {
  arr := []int{5, 2, 8, 1, 9, 4}
  second, err := FindSecondLargest(arr)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Второй по величине элемент:", second)
  }
}
