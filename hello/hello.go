package main

const englishHelloPrefix = "Hello, "

// Hello returns a greeting
func Hello(name string) string {
	if name == "" {
		name = "world!"
	}
	return englishHelloPrefix + name
}

func main() {
	println(Hello(""))
}
