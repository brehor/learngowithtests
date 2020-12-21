package main

const englishHelloPrefix = "Hello, "

// Hello returns a greeting
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	println(Hello(""))
}
