package main

import (
  "fmt"
  "github.com/patrickrodee/fila"
)

func main() {
  fmt.Printf("%+v\n", fila.Connect("https://foo.com"))
}