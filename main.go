package main

import "fmt"

//创建常量的价值是可以快速理解值的含义，有时还可以帮助提高性能。
const spanish = "Spanish"
const french = "French"
const englishPrefix = "hello "
const spanishPrefix = "hola "
const frenchPrefix = "bonjour "

const defaultName = "world"

func Hello(name string, language string) string {

	if len(name) == 0 {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	//返回中声明参数，它使你的代码更加清晰。
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

//但原则依然通用。TDD 是一门需要通过开发去实践的技能，通过将问题分解成更小的可测试的组件，你编写软件将会更加轻松。
func main() {

	fmt.Println(Hello("Luis", ""))
}
