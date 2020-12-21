package main

// Hello returns a greeting
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	println(Hello(""))
}
