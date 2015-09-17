package main

import (
  "fmt"
)

type Rectangle struct {
  length, width int
}

func (rect Rectangle) Area() int {
  return rect.length * rect.width
}
func main() {
  rect := Rectangle{length: 2, width: 4}
  fmt.Println(rect.Area())
}
