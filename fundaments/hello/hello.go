package main

import "fmt"


const enHelloPrefix = "Hello, "

func Hello(name string) string {
  if name == "" {
    return "Hello, World"
  }
  return enHelloPrefix + name
}


func main() {
  fmt.Println(Hello("World"))
}

