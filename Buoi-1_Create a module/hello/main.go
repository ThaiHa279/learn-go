package main

import "fmt"
import "example.com/greetings"

func main() {
	message := greetings.Hello("Thai Ha")
	fmt.Println(message)
}
