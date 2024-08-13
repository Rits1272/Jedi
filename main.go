package main

import (
  "fmt"
  "jedi/core"
)

func main() {
  fmt.Println("Welcome to Jedi - cache")

  core.Jedi.SetKey("Hello", "Jedi")
  var item *core.CacheItem = core.Jedi.GetKey("Hello")

  fmt.Println("GET KEY: HELLO ->", item.Value)
}
