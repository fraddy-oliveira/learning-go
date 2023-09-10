package main

import (
	"bufio"
	"fmt"
	"os"
)

var LANG = "GOLANG"

func main() {
	getNameFromInput()
}

func getNameFromInput() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	inputName, _ := reader.ReadString('\n')

	fmt.Printf("You have entered %s", inputName)
}

func variables() {
	// message := "Greetings from Go"

	another_message := "This is another message"

	fmt.Printf("Data type for another_message is %T \n", another_message)

	var defaultString string

	fmt.Printf("default string value is %s and const is %s \n", defaultString, LANG)
}
