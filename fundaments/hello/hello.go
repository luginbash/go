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

  return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
  switch lang {
  case french:
    prefix = frHelloPrefix
  case spanish:
    prefix = esHelloPrefix
  default:
    prefix = enHelloPrefix
  }
  return
}
