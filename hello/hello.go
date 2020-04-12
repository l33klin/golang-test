package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix = "你好, "


func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("Tom", ""))
}