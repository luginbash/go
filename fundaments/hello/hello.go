package main

const spanish = "Spanish"
const french = "French"
const enHelloPrefix = "Hello, "
const esHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
  if name == "" {
    name = "World"
  }
  if lang == spanish {
    return esHelloPrefix + name
  }

  if lang == french {
    return frHelloPrefix + name
  }
  return enHelloPrefix + name
}


