package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // все оргументы кроме имени программы
	
	if len(args) == 0 {
		printHelp()
		return
	}
	
	command := args[0]
	
	switch command {
	case "ping":
		ping(args[1:])
	case "hello":
		sayHello(args[1:])
	default:
		fmt.Println("Unknown command:", command)
		printHelp()
	}
}

//Выводим помощь
func printHelp() {
	fmt.Println("Usage:")
	fmt.Println(" go run main.go ping <something>")
	fmt.Println(" go run main.go hello <name>")
}

//Простуйшая функция ping
func ping(args []string){
	if len(args) == 0 {
		fmt.Println("ping what?")
		return
	}
	fmt.Println("Pong to", args[0])
}

//Функция hello
func sayHello(args []string) {
	if len(args) == 0 {
		fmt.Println("Hello, stranger!")
		return
	}
	fmt.Println("Hello,", args[0]+"!")
}