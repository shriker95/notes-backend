package main

import "fmt"

func Text() (msg string) {
	msg = "Hello World!"
	return
}

func main() {
	msg := Text()
	fmt.Println(msg)
}
