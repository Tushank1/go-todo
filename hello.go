package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func main() {
	var name string = "Tushank"
	greet(name)
}
